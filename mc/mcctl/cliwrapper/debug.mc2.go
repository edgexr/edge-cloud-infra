// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: debug.proto

package cliwrapper

import edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
import "strings"
import "github.com/mobiledgex/edge-cloud-infra/mc/ormapi"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/mobiledgex/edge-cloud/protogen"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

func (s *Client) EnableDebugLevels(uri, token string, in *ormapi.RegionDebugRequest) ([]edgeproto.DebugReply, int, error) {
	args := []string{"region", "EnableDebugLevels"}
	outlist := []edgeproto.DebugReply{}
	noconfig := strings.Split("Cmd", ",")
	ops := []runOp{
		withIgnore(noconfig),
	}
	st, err := s.runObjs(uri, token, args, in, &outlist, ops...)
	return outlist, st, err
}

func (s *Client) DisableDebugLevels(uri, token string, in *ormapi.RegionDebugRequest) ([]edgeproto.DebugReply, int, error) {
	args := []string{"region", "DisableDebugLevels"}
	outlist := []edgeproto.DebugReply{}
	noconfig := strings.Split("Cmd", ",")
	ops := []runOp{
		withIgnore(noconfig),
	}
	st, err := s.runObjs(uri, token, args, in, &outlist, ops...)
	return outlist, st, err
}

func (s *Client) ShowDebugLevels(uri, token string, in *ormapi.RegionDebugRequest) ([]edgeproto.DebugReply, int, error) {
	args := []string{"region", "ShowDebugLevels"}
	outlist := []edgeproto.DebugReply{}
	noconfig := strings.Split("Levels,Cmd", ",")
	ops := []runOp{
		withIgnore(noconfig),
	}
	st, err := s.runObjs(uri, token, args, in, &outlist, ops...)
	return outlist, st, err
}

func (s *Client) RunDebug(uri, token string, in *ormapi.RegionDebugRequest) ([]edgeproto.DebugReply, int, error) {
	args := []string{"region", "RunDebug"}
	outlist := []edgeproto.DebugReply{}
	noconfig := strings.Split("Levels", ",")
	ops := []runOp{
		withIgnore(noconfig),
	}
	st, err := s.runObjs(uri, token, args, in, &outlist, ops...)
	return outlist, st, err
}