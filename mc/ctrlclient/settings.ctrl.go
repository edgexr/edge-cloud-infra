// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: settings.proto

package ctrlclient

import (
	"context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/mobiledgex/edge-cloud-infra/mc/ormutil"
	edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
	"github.com/mobiledgex/edge-cloud/log"
	_ "github.com/mobiledgex/edge-cloud/protogen"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

func UpdateSettingsObj(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.Settings, connObj ClientConnMgr) (*edgeproto.Result, error) {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return nil, err
	}
	api := edgeproto.NewSettingsApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	return api.UpdateSettings(ctx, obj)
}

func ResetSettingsObj(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.Settings, connObj ClientConnMgr) (*edgeproto.Result, error) {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return nil, err
	}
	api := edgeproto.NewSettingsApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	return api.ResetSettings(ctx, obj)
}

func ShowSettingsObj(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.Settings, connObj ClientConnMgr) (*edgeproto.Settings, error) {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return nil, err
	}
	api := edgeproto.NewSettingsApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	return api.ShowSettings(ctx, obj)
}