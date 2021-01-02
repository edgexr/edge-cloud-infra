// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trustpolicy.proto

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

func CreateTrustPolicy(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionTrustPolicy{}
	success, err := ReadConn(c, &in)
	if !success {
		return err
	}
	defer CloseConn(c)
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("region", in.Region)
	log.SetTags(span, in.TrustPolicy.GetKey().GetTags())
	span.SetTag("org", in.TrustPolicy.Key.Organization)

	err = CreateTrustPolicyStream(ctx, rc, &in.TrustPolicy, func(res *edgeproto.Result) {
		payload := ormapi.StreamPayload{}
		payload.Data = res
		WriteStream(c, &payload)
	})
	if err != nil {
		WriteError(c, err)
	}
	return nil
}

func CreateTrustPolicyStream(ctx context.Context, rc *RegionContext, obj *edgeproto.TrustPolicy, cb func(res *edgeproto.Result)) error {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.Key.Organization,
			ResourceCloudlets, ActionManage, withRequiresOrg(obj.Key.Organization)); err != nil {
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
	api := edgeproto.NewTrustPolicyApiClient(rc.conn)
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
		cb(res)
	}
	return nil
}

func CreateTrustPolicyObj(ctx context.Context, rc *RegionContext, obj *edgeproto.TrustPolicy) ([]edgeproto.Result, error) {
	arr := []edgeproto.Result{}
	err := CreateTrustPolicyStream(ctx, rc, obj, func(res *edgeproto.Result) {
		arr = append(arr, *res)
	})
	return arr, err
}

func DeleteTrustPolicy(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionTrustPolicy{}
	success, err := ReadConn(c, &in)
	if !success {
		return err
	}
	defer CloseConn(c)
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("region", in.Region)
	log.SetTags(span, in.TrustPolicy.GetKey().GetTags())
	span.SetTag("org", in.TrustPolicy.Key.Organization)

	err = DeleteTrustPolicyStream(ctx, rc, &in.TrustPolicy, func(res *edgeproto.Result) {
		payload := ormapi.StreamPayload{}
		payload.Data = res
		WriteStream(c, &payload)
	})
	if err != nil {
		WriteError(c, err)
	}
	return nil
}

func DeleteTrustPolicyStream(ctx context.Context, rc *RegionContext, obj *edgeproto.TrustPolicy, cb func(res *edgeproto.Result)) error {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.Key.Organization,
			ResourceDeveloperPolicy, ActionManage); err != nil {
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
	api := edgeproto.NewTrustPolicyApiClient(rc.conn)
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
		cb(res)
	}
	return nil
}

func DeleteTrustPolicyObj(ctx context.Context, rc *RegionContext, obj *edgeproto.TrustPolicy) ([]edgeproto.Result, error) {
	arr := []edgeproto.Result{}
	err := DeleteTrustPolicyStream(ctx, rc, obj, func(res *edgeproto.Result) {
		arr = append(arr, *res)
	})
	return arr, err
}

func UpdateTrustPolicy(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionTrustPolicy{}
	success, err := ReadConn(c, &in)
	if !success {
		return err
	}
	defer CloseConn(c)
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("region", in.Region)
	log.SetTags(span, in.TrustPolicy.GetKey().GetTags())
	span.SetTag("org", in.TrustPolicy.Key.Organization)

	err = UpdateTrustPolicyStream(ctx, rc, &in.TrustPolicy, func(res *edgeproto.Result) {
		payload := ormapi.StreamPayload{}
		payload.Data = res
		WriteStream(c, &payload)
	})
	if err != nil {
		WriteError(c, err)
	}
	return nil
}

func UpdateTrustPolicyStream(ctx context.Context, rc *RegionContext, obj *edgeproto.TrustPolicy, cb func(res *edgeproto.Result)) error {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.Key.Organization,
			ResourceDeveloperPolicy, ActionManage); err != nil {
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
	api := edgeproto.NewTrustPolicyApiClient(rc.conn)
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
		cb(res)
	}
	return nil
}

func UpdateTrustPolicyObj(ctx context.Context, rc *RegionContext, obj *edgeproto.TrustPolicy) ([]edgeproto.Result, error) {
	arr := []edgeproto.Result{}
	err := UpdateTrustPolicyStream(ctx, rc, obj, func(res *edgeproto.Result) {
		arr = append(arr, *res)
	})
	return arr, err
}

func ShowTrustPolicy(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionTrustPolicy{}
	success, err := ReadConn(c, &in)
	if !success {
		return err
	}
	defer CloseConn(c)
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("region", in.Region)
	log.SetTags(span, in.TrustPolicy.GetKey().GetTags())
	span.SetTag("org", in.TrustPolicy.Key.Organization)

	err = ShowTrustPolicyStream(ctx, rc, &in.TrustPolicy, func(res *edgeproto.TrustPolicy) {
		payload := ormapi.StreamPayload{}
		payload.Data = res
		WriteStream(c, &payload)
	})
	if err != nil {
		WriteError(c, err)
	}
	return nil
}

func ShowTrustPolicyStream(ctx context.Context, rc *RegionContext, obj *edgeproto.TrustPolicy, cb func(res *edgeproto.TrustPolicy)) error {
	var authz *AuthzShow
	var err error
	if !rc.skipAuthz {
		authz, err = newShowAuthz(ctx, rc.region, rc.username, ResourceDeveloperPolicy, ActionView)
		if err == echo.ErrForbidden {
			return nil
		}
		if err != nil {
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
	api := edgeproto.NewTrustPolicyApiClient(rc.conn)
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
		if !rc.skipAuthz {
			if !authz.Ok(res.Key.Organization) {
				continue
			}
		}
		cb(res)
	}
	return nil
}

func ShowTrustPolicyObj(ctx context.Context, rc *RegionContext, obj *edgeproto.TrustPolicy) ([]edgeproto.TrustPolicy, error) {
	arr := []edgeproto.TrustPolicy{}
	err := ShowTrustPolicyStream(ctx, rc, obj, func(res *edgeproto.TrustPolicy) {
		arr = append(arr, *res)
	})
	return arr, err
}