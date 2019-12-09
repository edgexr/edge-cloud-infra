// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app_inst.proto

package testutil

import edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
import "os"
import "github.com/mobiledgex/edge-cloud-infra/mc/ormclient"
import "github.com/mobiledgex/edge-cloud-infra/mc/ormapi"
import "github.com/mobiledgex/edge-cloud/cli"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/mobiledgex/edge-cloud/protogen"
import _ "github.com/mobiledgex/edge-cloud/d-match-engine/dme-proto"
import _ "github.com/mobiledgex/edge-cloud/d-match-engine/dme-proto"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

func TestCreateAppInst(mcClient *ormclient.Client, uri, token, region string, in *edgeproto.AppInst) ([]edgeproto.Result, int, error) {
	dat := &ormapi.RegionAppInst{}
	dat.Region = region
	dat.AppInst = *in
	return mcClient.CreateAppInst(uri, token, dat)
}
func TestPermCreateAppInst(mcClient *ormclient.Client, uri, token, region, org string, targetCloudlet *edgeproto.CloudletKey) ([]edgeproto.Result, int, error) {
	in := &edgeproto.AppInst{}
	if targetCloudlet != nil {
		in.Key.ClusterInstKey.CloudletKey = *targetCloudlet
	}
	in.Key.AppKey.DeveloperKey.Name = org
	return TestCreateAppInst(mcClient, uri, token, region, in)
}

func TestDeleteAppInst(mcClient *ormclient.Client, uri, token, region string, in *edgeproto.AppInst) ([]edgeproto.Result, int, error) {
	dat := &ormapi.RegionAppInst{}
	dat.Region = region
	dat.AppInst = *in
	return mcClient.DeleteAppInst(uri, token, dat)
}
func TestPermDeleteAppInst(mcClient *ormclient.Client, uri, token, region, org string, targetCloudlet *edgeproto.CloudletKey) ([]edgeproto.Result, int, error) {
	in := &edgeproto.AppInst{}
	if targetCloudlet != nil {
		in.Key.ClusterInstKey.CloudletKey = *targetCloudlet
	}
	in.Key.AppKey.DeveloperKey.Name = org
	return TestDeleteAppInst(mcClient, uri, token, region, in)
}

func TestRefreshAppInst(mcClient *ormclient.Client, uri, token, region string, in *edgeproto.AppInst) ([]edgeproto.Result, int, error) {
	dat := &ormapi.RegionAppInst{}
	dat.Region = region
	dat.AppInst = *in
	return mcClient.RefreshAppInst(uri, token, dat)
}
func TestPermRefreshAppInst(mcClient *ormclient.Client, uri, token, region, org string, targetCloudlet *edgeproto.CloudletKey) ([]edgeproto.Result, int, error) {
	in := &edgeproto.AppInst{}
	if targetCloudlet != nil {
		in.Key.ClusterInstKey.CloudletKey = *targetCloudlet
	}
	in.Key.AppKey.DeveloperKey.Name = org
	return TestRefreshAppInst(mcClient, uri, token, region, in)
}

func TestUpdateAppInst(mcClient *ormclient.Client, uri, token, region string, in *edgeproto.AppInst) ([]edgeproto.Result, int, error) {
	dat := &ormapi.RegionAppInst{}
	dat.Region = region
	dat.AppInst = *in
	return mcClient.UpdateAppInst(uri, token, dat)
}
func TestPermUpdateAppInst(mcClient *ormclient.Client, uri, token, region, org string, targetCloudlet *edgeproto.CloudletKey) ([]edgeproto.Result, int, error) {
	in := &edgeproto.AppInst{}
	if targetCloudlet != nil {
		in.Key.ClusterInstKey.CloudletKey = *targetCloudlet
	}
	in.Key.AppKey.DeveloperKey.Name = org
	return TestUpdateAppInst(mcClient, uri, token, region, in)
}

func TestShowAppInst(mcClient *ormclient.Client, uri, token, region string, in *edgeproto.AppInst) ([]edgeproto.AppInst, int, error) {
	dat := &ormapi.RegionAppInst{}
	dat.Region = region
	dat.AppInst = *in
	return mcClient.ShowAppInst(uri, token, dat)
}
func TestPermShowAppInst(mcClient *ormclient.Client, uri, token, region, org string) ([]edgeproto.AppInst, int, error) {
	in := &edgeproto.AppInst{}
	in.Key.AppKey.DeveloperKey.Name = org
	return TestShowAppInst(mcClient, uri, token, region, in)
}

func RunMcAppInstApi(mcClient ormclient.Api, uri, token, region string, data *[]edgeproto.AppInst, dataIn interface{}, rc *bool, mode string) {
	var dataInList []interface{}
	var ok bool
	if dataIn != nil {
		dataInList, ok = dataIn.([]interface{})
		if !ok {
			fmt.Fprintf(os.Stderr, "invalid data in appInst: %v\n", dataIn)
			os.Exit(1)
		}
	}
	for ii, appInst := range *data {
		dataMap, ok := dataInList[ii].(map[string]interface{})
		if !ok {
			fmt.Fprintf(os.Stderr, "invalid data in appInst: %v\n", dataInList[ii])
			os.Exit(1)
		}
		in := &ormapi.RegionAppInst{
			Region:  region,
			AppInst: appInst,
		}
		switch mode {
		case "create":
			_, st, err := mcClient.CreateAppInst(uri, token, in)
			checkMcErr("CreateAppInst", st, err, rc)
		case "delete":
			_, st, err := mcClient.DeleteAppInst(uri, token, in)
			checkMcErr("DeleteAppInst", st, err, rc)
		case "update":
			in.AppInst.Fields = cli.GetSpecifiedFields(dataMap, &in.AppInst, cli.YamlNamespace)
			_, st, err := mcClient.UpdateAppInst(uri, token, in)
			checkMcErr("UpdateAppInst", st, err, rc)
		default:
			return
		}
	}
}