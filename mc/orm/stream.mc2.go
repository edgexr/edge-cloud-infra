// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stream.proto

package orm

import (
	"context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/labstack/echo"
	"github.com/mobiledgex/edge-cloud-infra/mc/ormapi"
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

func StreamAppInst(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionAppInstKey{}
	success, err := ReadConn(c, &in)
	if !success {
		return err
	}
	defer CloseConn(c)
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("region", in.Region)
	span.SetTag("org", in.AppInstKey.AppKey.Organization)

	err = StreamAppInstStream(ctx, rc, &in.AppInstKey, func(res *edgeproto.StreamMsg) {
		payload := ormapi.StreamPayload{}
		payload.Data = res
		WriteStream(c, &payload)
	})
	if err != nil {
		WriteError(c, err)
	}
	return nil
}

func StreamAppInstStream(ctx context.Context, rc *RegionContext, obj *edgeproto.AppInstKey, cb func(res *edgeproto.StreamMsg)) error {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.AppKey.Organization,
			ResourceAppInsts, ActionView); err != nil {
			return err
		}
	}
	if rc.conn == nil {
		conn, err := connectController(ctx, rc.region)
		if err != nil {
			return err
		}
		rc.conn = conn
		defer func() {
			rc.conn.Close()
			rc.conn = nil
		}()
	}
	api := edgeproto.NewStreamObjApiClient(rc.conn)
	stream, err := api.StreamAppInst(ctx, obj)
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
		cb(res)
	}
	return nil
}

func StreamAppInstObj(ctx context.Context, rc *RegionContext, obj *edgeproto.AppInstKey) ([]edgeproto.StreamMsg, error) {
	arr := []edgeproto.StreamMsg{}
	err := StreamAppInstStream(ctx, rc, obj, func(res *edgeproto.StreamMsg) {
		arr = append(arr, *res)
	})
	return arr, err
}

func StreamClusterInst(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionClusterInstKey{}
	success, err := ReadConn(c, &in)
	if !success {
		return err
	}
	defer CloseConn(c)
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("region", in.Region)
	span.SetTag("org", in.ClusterInstKey.Organization)

	err = StreamClusterInstStream(ctx, rc, &in.ClusterInstKey, func(res *edgeproto.StreamMsg) {
		payload := ormapi.StreamPayload{}
		payload.Data = res
		WriteStream(c, &payload)
	})
	if err != nil {
		WriteError(c, err)
	}
	return nil
}

func StreamClusterInstStream(ctx context.Context, rc *RegionContext, obj *edgeproto.ClusterInstKey, cb func(res *edgeproto.StreamMsg)) error {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.Organization,
			ResourceClusterInsts, ActionView); err != nil {
			return err
		}
	}
	if rc.conn == nil {
		conn, err := connectController(ctx, rc.region)
		if err != nil {
			return err
		}
		rc.conn = conn
		defer func() {
			rc.conn.Close()
			rc.conn = nil
		}()
	}
	api := edgeproto.NewStreamObjApiClient(rc.conn)
	stream, err := api.StreamClusterInst(ctx, obj)
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
		cb(res)
	}
	return nil
}

func StreamClusterInstObj(ctx context.Context, rc *RegionContext, obj *edgeproto.ClusterInstKey) ([]edgeproto.StreamMsg, error) {
	arr := []edgeproto.StreamMsg{}
	err := StreamClusterInstStream(ctx, rc, obj, func(res *edgeproto.StreamMsg) {
		arr = append(arr, *res)
	})
	return arr, err
}

func StreamCloudlet(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionCloudletKey{}
	success, err := ReadConn(c, &in)
	if !success {
		return err
	}
	defer CloseConn(c)
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("region", in.Region)
	span.SetTag("org", in.CloudletKey.Organization)

	err = StreamCloudletStream(ctx, rc, &in.CloudletKey, func(res *edgeproto.StreamMsg) {
		payload := ormapi.StreamPayload{}
		payload.Data = res
		WriteStream(c, &payload)
	})
	if err != nil {
		WriteError(c, err)
	}
	return nil
}

func StreamCloudletStream(ctx context.Context, rc *RegionContext, obj *edgeproto.CloudletKey, cb func(res *edgeproto.StreamMsg)) error {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.Organization,
			ResourceCloudlets, ActionView); err != nil {
			return err
		}
	}
	if rc.conn == nil {
		conn, err := connectController(ctx, rc.region)
		if err != nil {
			return err
		}
		rc.conn = conn
		defer func() {
			rc.conn.Close()
			rc.conn = nil
		}()
	}
	api := edgeproto.NewStreamObjApiClient(rc.conn)
	stream, err := api.StreamCloudlet(ctx, obj)
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
		cb(res)
	}
	return nil
}

func StreamCloudletObj(ctx context.Context, rc *RegionContext, obj *edgeproto.CloudletKey) ([]edgeproto.StreamMsg, error) {
	arr := []edgeproto.StreamMsg{}
	err := StreamCloudletStream(ctx, rc, obj, func(res *edgeproto.StreamMsg) {
		arr = append(arr, *res)
	})
	return arr, err
}