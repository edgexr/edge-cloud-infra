// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trustpolicy.proto

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
	"io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

func CreateTrustPolicyStream(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.TrustPolicy, connObj ClientConnMgr, cb func(res *edgeproto.Result) error) error {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return err
	}
	api := edgeproto.NewTrustPolicyApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	stream, err := api.CreateTrustPolicy(ctx, obj)
	if err != nil {
		return err
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return err
		}
		err = cb(res)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteTrustPolicyStream(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.TrustPolicy, connObj ClientConnMgr, cb func(res *edgeproto.Result) error) error {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return err
	}
	api := edgeproto.NewTrustPolicyApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	stream, err := api.DeleteTrustPolicy(ctx, obj)
	if err != nil {
		return err
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return err
		}
		err = cb(res)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateTrustPolicyStream(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.TrustPolicy, connObj ClientConnMgr, cb func(res *edgeproto.Result) error) error {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return err
	}
	api := edgeproto.NewTrustPolicyApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	stream, err := api.UpdateTrustPolicy(ctx, obj)
	if err != nil {
		return err
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return err
		}
		err = cb(res)
		if err != nil {
			return err
		}
	}
	return nil
}

type ShowTrustPolicyAuthz interface {
	Ok(obj *edgeproto.TrustPolicy) (bool, bool)
	Filter(obj *edgeproto.TrustPolicy)
}

func ShowTrustPolicyStream(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.TrustPolicy, connObj ClientConnMgr, authz ShowTrustPolicyAuthz, cb func(res *edgeproto.TrustPolicy) error) error {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return err
	}
	api := edgeproto.NewTrustPolicyApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	stream, err := api.ShowTrustPolicy(ctx, obj)
	if err != nil {
		return err
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return err
		}
		if !rc.SkipAuthz {
			if authz != nil {
				authzOk, filterOutput := authz.Ok(res)
				if !authzOk {
					continue
				}
				if filterOutput {
					authz.Filter(res)
				}
			}
		}
		err = cb(res)
		if err != nil {
			return err
		}
	}
	return nil
}
