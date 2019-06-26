package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	platform "github.com/mobiledgex/edge-cloud-infra/shepherd/shepherd_platform"
	"github.com/mobiledgex/edge-cloud-infra/shepherd/shepherd_platform/shepherd_dind"
	"github.com/mobiledgex/edge-cloud-infra/shepherd/shepherd_platform/shepherd_openstack"
	"github.com/mobiledgex/edge-cloud/cloud-resource-manager/k8smgmt"
	"github.com/mobiledgex/edge-cloud/cloudcommon"
	influxq "github.com/mobiledgex/edge-cloud/controller/influxq_client"
	"github.com/mobiledgex/edge-cloud/edgeproto"
	"github.com/mobiledgex/edge-cloud/log"
	"github.com/mobiledgex/edge-cloud/notify"
)

var influxdb = flag.String("influxdb", "http://0.0.0.0:8086", "InfluxDB address to export to")
var debugLevels = flag.String("d", "", fmt.Sprintf("comma separated list of %v", log.DebugLevelStrings))
var notifyAddrs = flag.String("notifyAddrs", "127.0.0.1:51001", "CRM notify listener addresses")
var tlsCertFile = flag.String("tls", "", "server tls cert file.  Keyfile and CA file mex-ca.crt must be in same directory")
var collectInterval = flag.Duration("interval", time.Second*15, "Metrics collection interval")
var platformName = flag.String("platform", "", "Platform type of Cloudlet")
var vaultAddr = flag.String("vaultAddr", "", "Address to vault")
var physicalName = flag.String("physicalName", "", "Physical infrastructure cloudlet name, defaults to cloudlet name in cloudletKey")
var cloudletKeyStr = flag.String("cloudletKey", "", "Json or Yaml formatted cloudletKey for the cloudlet in which this CRM is instantiated; e.g. '{\"operator_key\":{\"name\":\"DMUUS\"},\"name\":\"tmocloud1\"}'")

var promQCpuClust = "sum(rate(container_cpu_usage_seconds_total%7Bid%3D%22%2F%22%7D%5B1m%5D))%2Fsum(machine_cpu_cores)*100"
var promQMemClust = "sum(container_memory_working_set_bytes%7Bid%3D%22%2F%22%7D)%2Fsum(machine_memory_bytes)*100"
var promQDiskClust = "sum(container_fs_usage_bytes%7Bdevice%3D~%22%5E%2Fdev%2F%5Bsv%5Dd%5Ba-z%5D%5B1-9%5D%24%22%2Cid%3D%22%2F%22%7D)%2Fsum(container_fs_limit_bytes%7Bdevice%3D~%22%5E%2Fdev%2F%5Bsv%5Dd%5Ba-z%5D%5B1-9%5D%24%22%2Cid%3D%22%2F%22%7D)*100"
var promQSendBytesRateClust = "sum(irate(container_network_transmit_bytes_total%5B1m%5D))"
var promQRecvBytesRateClust = "sum(irate(container_network_receive_bytes_total%5B1m%5D))"
var promQTcpConnClust = "node_netstat_Tcp_CurrEstab"
var promQTcpRetransClust = "node_netstat_Tcp_RetransSegs"
var promQUdpSendPktsClust = "node_netstat_Udp_OutDatagrams"
var promQUdpRecvPktsClust = "node_netstat_Udp_InDatagrams"
var promQUdpRecvErr = "node_netstat_Udp_InErrors"

var promQCpuPod = "sum(rate(container_cpu_usage_seconds_total%7Bimage!%3D%22%22%7D%5B1m%5D))by(pod_name)"
var promQMemPod = "sum(container_memory_working_set_bytes%7Bimage!%3D%22%22%7D)by(pod_name)"
var promQNetRecvRate = "sum(irate(container_network_receive_bytes_total%7Bimage!%3D%22%22%7D%5B1m%5D))by(pod_name)"
var promQNetSendRate = "sum(irate(container_network_transmit_bytes_total%7Bimage!%3D%22%22%7D%5B1m%5D))by(pod_name)"

var Env = map[string]string{
	"INFLUXDB_USER": "root",
	"INFLUXDB_PASS": "root",
}

//map keeping track of all the currently running prometheuses
var promMap map[string]*PromStats
var MEXPrometheusAppName = cloudcommon.MEXPrometheusAppName
var AppInstCache edgeproto.AppInstCache
var ClusterInstCache edgeproto.ClusterInstCache

var cloudletKey edgeproto.CloudletKey
var pf platform.Platform

var InfluxDBName = cloudcommon.DeveloperMetricsDbName
var influxQ *influxq.InfluxQ

var sigChan chan os.Signal

func appInstCb(old *edgeproto.AppInst, new *edgeproto.AppInst) {
	//check for prometheus
	if new.Key.AppKey.Name != MEXPrometheusAppName {
		return
	}
	var mapKey = k8smgmt.GetK8sNodeNameSuffix(&new.Key.ClusterInstKey)
	stats, exists := promMap[mapKey]
	if new.State == edgeproto.TrackedState_READY {
		log.DebugLog(log.DebugLevelMetrics, "New Prometheus instance detected", "clustername", mapKey)
		//get address of prometheus.
		clusterInst := edgeproto.ClusterInst{}
		found := ClusterInstCache.Get(&new.Key.ClusterInstKey, &clusterInst)
		if !found {
			log.DebugLog(log.DebugLevelMetrics, "Unable to find clusterInst for prometheus")
			return
		}
		clustIP, err := pf.GetClusterIP(&clusterInst)
		if err != nil {
			log.DebugLog(log.DebugLevelMetrics, "error getting clusterIP", "err", err.Error())
		}
		port := new.MappedPorts[0].PublicPort
		promAddress := fmt.Sprintf("%s:%d", clustIP, port)
		log.DebugLog(log.DebugLevelMetrics, "prometheus found", "promAddress", promAddress)
		if !exists {
			stats = NewPromStats(promAddress, *collectInterval, sendMetric, &clusterInst, pf)
			promMap[mapKey] = stats
			stats.Start()
		} else { //somehow this cluster's prometheus was already registered
			log.DebugLog(log.DebugLevelMetrics, "Error, Prometheus app already registered for this cluster")
		}
	} else { //if its anything other than ready just stop it
		//try to remove it from the prommap
		if exists {
			delete(promMap, mapKey)
			stats.Stop()
		}
	}
}

func getPlatform() (platform.Platform, error) {
	var plat platform.Platform
	var err error
	switch *platformName {
	case "dind":
		plat = &shepherd_dind.Platform{}
	case "openstack":
		plat = &shepherd_openstack.Platform{}
	default:
		err = fmt.Errorf("Platform %s not supported", *platformName)
	}
	return plat, err
}

func main() {
	flag.Parse()
	log.SetDebugLevelStrs(*debugLevels)

	cloudcommon.ParseMyCloudletKey(false, cloudletKeyStr, &cloudletKey)
	log.DebugLog(log.DebugLevelMetrics, "InfluxDB found", "influxdb addr", influxdb)
	log.DebugLog(log.DebugLevelMetrics, "Metrics collection", "interval", collectInterval)
	var err error
	pf, err = getPlatform()
	if err != nil {
		log.FatalLog("Failed to get platform", "platformName", platformName, "err", err)
	}
	pf.Init(&cloudletKey, *physicalName, *vaultAddr)
	influxQ = influxq.NewInfluxQ(InfluxDBName)
	err = influxQ.Start(*influxdb, "")
	if err != nil {
		log.FatalLog("Failed to start influx queue",
			"address", *influxdb, "err", err)
	}
	defer influxQ.Stop()

	promMap = make(map[string]*PromStats)

	//register shepherd to receive appinst notifications from crm
	edgeproto.InitAppInstCache(&AppInstCache)
	AppInstCache.SetUpdatedCb(appInstCb)
	edgeproto.InitClusterInstCache(&ClusterInstCache)
	addrs := strings.Split(*notifyAddrs, ",")
	notifyClient := notify.NewClient(addrs, *tlsCertFile)
	notifyClient.RegisterRecvAppInstCache(&AppInstCache)
	notifyClient.RegisterRecvClusterInstCache(&ClusterInstCache)
	notifyClient.Start()
	defer notifyClient.Stop()

	sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	log.DebugLog(log.DebugLevelMetrics, "Ready")

	// wait until process in killed/interrupted
	sig := <-sigChan
	fmt.Println(sig)
}

func sendMetric(metric *edgeproto.Metric) {
	influxQ.AddMetric(metric)
}