// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: appinst.proto

package ormclient

import edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
import "github.com/mobiledgex/edge-cloud-infra/mc/ormapi"
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

func (s *Client) CreateAppInst(uri, token string, in *ormapi.RegionAppInst) ([]edgeproto.Result, int, error) {
	out := edgeproto.Result{}
	outlist := []edgeproto.Result{}
	status, err := s.PostJsonStreamOut(uri+"/auth/ctrl/CreateAppInst", token, in, &out, func() {
		outlist = append(outlist, out)
	})
	return outlist, status, err
}

func (s *Client) DeleteAppInst(uri, token string, in *ormapi.RegionAppInst) ([]edgeproto.Result, int, error) {
	out := edgeproto.Result{}
	outlist := []edgeproto.Result{}
	status, err := s.PostJsonStreamOut(uri+"/auth/ctrl/DeleteAppInst", token, in, &out, func() {
		outlist = append(outlist, out)
	})
	return outlist, status, err
}

func (s *Client) RefreshAppInst(uri, token string, in *ormapi.RegionAppInst) ([]edgeproto.Result, int, error) {
	out := edgeproto.Result{}
	outlist := []edgeproto.Result{}
	status, err := s.PostJsonStreamOut(uri+"/auth/ctrl/RefreshAppInst", token, in, &out, func() {
		outlist = append(outlist, out)
	})
	return outlist, status, err
}

func (s *Client) UpdateAppInst(uri, token string, in *ormapi.RegionAppInst) ([]edgeproto.Result, int, error) {
	out := edgeproto.Result{}
	outlist := []edgeproto.Result{}
	status, err := s.PostJsonStreamOut(uri+"/auth/ctrl/UpdateAppInst", token, in, &out, func() {
		outlist = append(outlist, out)
	})
	return outlist, status, err
}

func (s *Client) ShowAppInst(uri, token string, in *ormapi.RegionAppInst) ([]edgeproto.AppInst, int, error) {
	out := edgeproto.AppInst{}
	outlist := []edgeproto.AppInst{}
	status, err := s.PostJsonStreamOut(uri+"/auth/ctrl/ShowAppInst", token, in, &out, func() {
		outlist = append(outlist, out)
	})
	return outlist, status, err
}

type AppInstApiClient interface {
	CreateAppInst(uri, token string, in *ormapi.RegionAppInst) ([]edgeproto.Result, int, error)
	DeleteAppInst(uri, token string, in *ormapi.RegionAppInst) ([]edgeproto.Result, int, error)
	RefreshAppInst(uri, token string, in *ormapi.RegionAppInst) ([]edgeproto.Result, int, error)
	UpdateAppInst(uri, token string, in *ormapi.RegionAppInst) ([]edgeproto.Result, int, error)
	ShowAppInst(uri, token string, in *ormapi.RegionAppInst) ([]edgeproto.AppInst, int, error)
}