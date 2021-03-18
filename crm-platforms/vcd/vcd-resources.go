package vcd

import (
	"context"
	"fmt"
	"strings"

	"github.com/mobiledgex/edge-cloud-infra/infracommon"
	"github.com/mobiledgex/edge-cloud-infra/vmlayer"

	"github.com/mobiledgex/edge-cloud/cloudcommon"
	"github.com/mobiledgex/edge-cloud/edgeproto"
	"github.com/mobiledgex/edge-cloud/log"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	vcdtypes "github.com/vmware/go-vcloud-director/v2/types/v56"
)

var (
	ResourceInstances   = "Instances"
	ResourceExternalIps = "External IPs"
)

type VcdResources struct {
	VmsUsed         uint64
	ExternalIpsUsed uint64
}

func (v *VcdPlatform) GetPlatformResourceInfo(ctx context.Context) (*vmlayer.PlatformResources, error) {
	log.SpanLog(ctx, log.DebugLevelInfra, "GetPlatformResourceInfo")

	var resources = vmlayer.PlatformResources{}
	resinfo, err := v.GetCloudletInfraResourcesInfo(ctx)
	if err != nil {
		return nil, err
	}
	// TODO, 2 similar structs for the same concept should be revisited
	for _, r := range resinfo {
		switch r.Name {
		case cloudcommon.ResourceVcpus:
			resources.VCpuMax = r.InfraMaxValue
			resources.VCpuUsed = r.Value
		case cloudcommon.ResourceRamMb:
			resources.MemMax = r.InfraMaxValue
			resources.MemUsed = r.Value
		case ResourceExternalIps:
			resources.Ipv4Max = r.InfraMaxValue
			resources.Ipv4Used = r.Value
		}
	}
	return &resources, nil
}

func (v *VcdPlatform) GetCloudletInfraResourcesInfo(ctx context.Context) ([]edgeproto.InfraResource, error) {
	log.SpanLog(ctx, log.DebugLevelInfra, "GetCloudletInfraResourcesInfo")
	resInfo := []edgeproto.InfraResource{}
	vcdClient := v.GetVcdClientFromContext(ctx)

	if vcdClient == nil {
		return nil, fmt.Errorf(NoVCDClientInContext)
	}
	vdc, err := v.GetVdc(ctx, vcdClient)
	if err != nil {
		return nil, err
	}

	org, err := v.GetOrg(ctx, vcdClient)
	if err != nil {
		return nil, fmt.Errorf("Error getting VDC Org: %v", err)
	}

	// get the cpu speed to calculate number of VMs used.  When we create VMs we specify the number of VCPUs, but
	// to find the quotas and numbers used, we have to search for the CPU speed and calculate.
	cpuSpeed := v.GetVcpuSpeedOverride(ctx)
	if cpuSpeed == 0 {
		// retrieve from admin org
		adminOrg, err := govcd.GetAdminOrgByName(vcdClient, org.Org.Name)
		if err != nil {
			log.SpanLog(ctx, log.DebugLevelInfra, "Unable to get AdminOrg", "orgName", "org.Org.Name", "error", err)
			return nil, fmt.Errorf("Unable to get AdminOrg named: %s - %v", org.Org.Name, err)
		}
		adminVdc, err := adminOrg.GetAdminVdcByName(v.GetVDCName())
		if err != nil {
			log.SpanLog(ctx, log.DebugLevelInfra, "Unable to get AdminVdc", "adminOrgName", "adminOrg.Org.Name", "error", err)
			return nil, fmt.Errorf("Unable to get AdminVcd named: %s - %v", v.GetVDCName(), err)
		}
		// VMW stores the speed in 2 different places, the first of which is generally nil in our testing
		if adminVdc.AdminVdc.VCpuInMhz != nil && *adminVdc.AdminVdc.VCpuInMhz != 0 {
			cpuSpeed = *adminVdc.AdminVdc.VCpuInMhz
			log.SpanLog(ctx, log.DebugLevelInfra, "Using cpu speed from admin VCpuInMhz", "cpuSpeed", cpuSpeed)
		} else {
			if adminVdc.AdminVdc.VCpuInMhz2 != nil && *adminVdc.AdminVdc.VCpuInMhz2 != 0 {
				cpuSpeed = *adminVdc.AdminVdc.VCpuInMhz2
				log.SpanLog(ctx, log.DebugLevelInfra, "Using cpu speed from admin VCpuInMhz2", "cpuSpeed", cpuSpeed)
			} else {
				return nil, fmt.Errorf("No cpu speed in organization")
			}
		}
	} else {
		log.SpanLog(ctx, log.DebugLevelInfra, "Using cpu speed from properties", "cpuSpeed", cpuSpeed)
	}

	vmlist, err := vcdClient.Client.QueryVmList(vcdtypes.VmQueryFilterAll)
	if err != nil {
		return nil, fmt.Errorf("Failed to query VmList: %v", err)
	}
	extNet, err := v.GetExtNetwork(ctx, vcdClient)
	if err != nil {
		return nil, err
	}
	ipScopes := extNet.OrgVDCNetwork.Configuration.IPScopes.IPScope
	ranges := []string{}
	for _, ips := range ipScopes {
		mask, err := MaskToCidr(ips.Netmask)
		if err != nil {
			return nil, fmt.Errorf("MaskToCidr failed - %s - %v", ips.Netmask, err)
		}
		for _, ipr := range ips.IPRanges.IPRange {
			ranges = append(ranges, fmt.Sprintf("%s/%s-%s/%s", ipr.StartAddress, mask, ipr.EndAddress, mask))
		}
	}
	iprangeString := strings.Join(ranges, ",")
	availIps, err := infracommon.ParseIpRanges(iprangeString)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse ip ranges from org vcd network ranges: %s - %v", iprangeString, err)
	}
	var usedIps uint64 = 0
	extNetName := v.vmProperties.GetCloudletExternalNetwork()
	for _, vm := range vmlist {
		if vm.NetworkName == extNetName {
			usedIps++
		}
	}

	resInfo = append(resInfo, edgeproto.InfraResource{
		Name:          ResourceExternalIps,
		InfraMaxValue: uint64(len(availIps)),
		Value:         usedIps,
	})
	resInfo = append(resInfo, edgeproto.InfraResource{
		Name:          ResourceInstances,
		InfraMaxValue: uint64(vdc.Vdc.VMQuota),
		Value:         uint64(len(vmlist)),
	})
	for _, cap := range vdc.Vdc.ComputeCapacity {
		resInfo = append(resInfo, edgeproto.InfraResource{
			Name:          cloudcommon.ResourceVcpus,
			Value:         uint64((cap.CPU.Used) / cpuSpeed),
			InfraMaxValue: uint64((cap.CPU.Limit) / cpuSpeed),
		})
		resInfo = append(resInfo, edgeproto.InfraResource{
			Name:          cloudcommon.ResourceRamMb,
			Value:         uint64(cap.Memory.Used),
			InfraMaxValue: uint64(cap.Memory.Limit),
		})
	}
	return resInfo, nil
}

func (v *VcdPlatform) GetCloudletResourceQuotaProps(ctx context.Context) (*edgeproto.CloudletResourceQuotaProps, error) {
	log.SpanLog(ctx, log.DebugLevelInfra, "GetCloudletResourceQuotaProps")

	return &edgeproto.CloudletResourceQuotaProps{
		Properties: []edgeproto.InfraResource{
			{
				Name:        ResourceExternalIps,
				Description: "Limit on how many external IPs are available",
			},
			{
				Name:        ResourceInstances,
				Description: "Limit on number of instances that can be provisioned",
			},
		},
	}, nil

}

func getVcdResources(ctx context.Context, cloudlet *edgeproto.Cloudlet, resources []edgeproto.VMResource) *VcdResources {
	log.SpanLog(ctx, log.DebugLevelInfra, "getVcdResources")
	var vRes VcdResources
	for _, vmRes := range resources {
		log.SpanLog(ctx, log.DebugLevelInfra, "getVcdResources", "vmRes", vmRes)

		// Number of Instances = Number of resources
		vRes.VmsUsed += 1
		if vmRes.Type == cloudcommon.VMTypeRootLB {
			vRes.ExternalIpsUsed += 1
		}
	}
	return &vRes
}

// called by controller, make sure it doesn't make any calls to infra API
func (v *VcdPlatform) GetClusterAdditionalResources(ctx context.Context, cloudlet *edgeproto.Cloudlet, vmResources []edgeproto.VMResource, infraResMap map[string]edgeproto.InfraResource) map[string]edgeproto.InfraResource {
	// resource name -> resource units
	cloudletRes := map[string]string{
		ResourceInstances:   "",
		ResourceExternalIps: "",
	}
	resInfo := make(map[string]edgeproto.InfraResource)
	for resName, resUnits := range cloudletRes {
		resMax := uint64(0)
		if infraRes, ok := infraResMap[resName]; ok {
			resMax = infraRes.InfraMaxValue
		}
		resInfo[resName] = edgeproto.InfraResource{
			Name:          resName,
			InfraMaxValue: resMax,
			Units:         resUnits,
		}
	}
	vRes := getVcdResources(ctx, cloudlet, vmResources)
	outInfo, ok := resInfo[ResourceInstances]
	if ok {
		outInfo.Value += vRes.VmsUsed
		resInfo[ResourceInstances] = outInfo
	}
	outInfo, ok = resInfo[ResourceExternalIps]
	if ok {
		outInfo.Value += vRes.ExternalIpsUsed
		resInfo[ResourceExternalIps] = outInfo
	}
	return resInfo
}

func (v *VcdPlatform) GetClusterAdditionalResourceMetric(ctx context.Context, cloudlet *edgeproto.Cloudlet, resMetric *edgeproto.Metric, resources []edgeproto.VMResource) error {
	log.SpanLog(ctx, log.DebugLevelInfra, "GetClusterAdditionalResourceMetric ")
	vRes := getVcdResources(ctx, cloudlet, resources)
	resMetric.AddIntVal("VMsUsed", vRes.VmsUsed)
	resMetric.AddIntVal("externalIpsUsed", vRes.ExternalIpsUsed)

	return nil
}