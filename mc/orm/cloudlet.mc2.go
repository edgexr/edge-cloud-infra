// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cloudlet.proto

package orm

import edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
import "github.com/labstack/echo"
import "context"
import "io"
import "github.com/mobiledgex/edge-cloud/log"
import "github.com/mobiledgex/edge-cloud-infra/mc/ormapi"
import "google.golang.org/grpc/status"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/mobiledgex/edge-cloud/protogen"
import _ "github.com/mobiledgex/edge-cloud/d-match-engine/dme-proto"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

var streamCloudlet = &StreamObj{}

func StreamCloudlet(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionCloudlet{}
	success, err := ReadConn(c, &in)
	if !success {
		return err
	}
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("org", in.Cloudlet.Key.Organization)

	streamer := streamCloudlet.Get(in.Cloudlet.Key)
	if streamer != nil {
		payload := ormapi.StreamPayload{}
		streamCh := streamer.Subscribe()
		serverClosed := make(chan bool)
		go func() {
			for streamMsg := range streamCh {
				switch out := streamMsg.(type) {
				case string:
					payload.Data = &edgeproto.Result{Message: out}
					WriteStream(c, &payload)
				case error:
					WriteError(c, out)
				default:
					WriteError(c, fmt.Errorf("Unsupported message type received: %v", streamMsg))
				}
			}
			CloseConn(c)
			serverClosed <- true
		}()
		// Wait for client/server to close
		// * Server closure is set via above serverClosed flag
		// * Client closure is sent from client via a message
		WaitForConnClose(c, serverClosed)
		streamer.Unsubscribe(streamCh)
	} else {
		WriteError(c, fmt.Errorf("Key doesn't exist"))
		CloseConn(c)
	}
	return nil
}

func CreateCloudlet(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionCloudlet{}
	success, err := ReadConn(c, &in)
	if !success {
		return err
	}
	defer CloseConn(c)
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("org", in.Cloudlet.Key.Organization)

	streamer := NewStreamer()
	defer streamer.Stop()
	streamAdded := false

	err = CreateCloudletStream(ctx, rc, &in.Cloudlet, func(res *edgeproto.Result) {
		if !streamAdded {
			streamCloudlet.Add(in.Cloudlet.Key, streamer)
			streamAdded = true
		}
		payload := ormapi.StreamPayload{}
		payload.Data = res
		streamer.Publish(res.Message)
		WriteStream(c, &payload)
	})
	if err != nil {
		streamer.Publish(err)
		WriteError(c, err)
	}
	if streamAdded {
		streamCloudlet.Remove(in.Cloudlet.Key, streamer)
	}
	return nil
}

func CreateCloudletStream(ctx context.Context, rc *RegionContext, obj *edgeproto.Cloudlet, cb func(res *edgeproto.Result)) error {
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
	api := edgeproto.NewCloudletApiClient(rc.conn)
	stream, err := api.CreateCloudlet(ctx, obj)
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

func CreateCloudletObj(ctx context.Context, rc *RegionContext, obj *edgeproto.Cloudlet) ([]edgeproto.Result, error) {
	arr := []edgeproto.Result{}
	err := CreateCloudletStream(ctx, rc, obj, func(res *edgeproto.Result) {
		arr = append(arr, *res)
	})
	return arr, err
}

func DeleteCloudlet(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionCloudlet{}
	success, err := ReadConn(c, &in)
	if !success {
		return err
	}
	defer CloseConn(c)
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("org", in.Cloudlet.Key.Organization)

	streamer := NewStreamer()
	defer streamer.Stop()
	streamAdded := false

	err = DeleteCloudletStream(ctx, rc, &in.Cloudlet, func(res *edgeproto.Result) {
		if !streamAdded {
			streamCloudlet.Add(in.Cloudlet.Key, streamer)
			streamAdded = true
		}
		payload := ormapi.StreamPayload{}
		payload.Data = res
		streamer.Publish(res.Message)
		WriteStream(c, &payload)
	})
	if err != nil {
		streamer.Publish(err)
		WriteError(c, err)
	}
	if streamAdded {
		streamCloudlet.Remove(in.Cloudlet.Key, streamer)
	}
	return nil
}

func DeleteCloudletStream(ctx context.Context, rc *RegionContext, obj *edgeproto.Cloudlet, cb func(res *edgeproto.Result)) error {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.Key.Organization,
			ResourceCloudlets, ActionManage); err != nil {
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
	api := edgeproto.NewCloudletApiClient(rc.conn)
	stream, err := api.DeleteCloudlet(ctx, obj)
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

func DeleteCloudletObj(ctx context.Context, rc *RegionContext, obj *edgeproto.Cloudlet) ([]edgeproto.Result, error) {
	arr := []edgeproto.Result{}
	err := DeleteCloudletStream(ctx, rc, obj, func(res *edgeproto.Result) {
		arr = append(arr, *res)
	})
	return arr, err
}

func UpdateCloudlet(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionCloudlet{}
	success, err := ReadConn(c, &in)
	if !success {
		return err
	}
	defer CloseConn(c)
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("org", in.Cloudlet.Key.Organization)

	streamer := NewStreamer()
	defer streamer.Stop()
	streamAdded := false

	err = UpdateCloudletStream(ctx, rc, &in.Cloudlet, func(res *edgeproto.Result) {
		if !streamAdded {
			streamCloudlet.Add(in.Cloudlet.Key, streamer)
			streamAdded = true
		}
		payload := ormapi.StreamPayload{}
		payload.Data = res
		streamer.Publish(res.Message)
		WriteStream(c, &payload)
	})
	if err != nil {
		streamer.Publish(err)
		WriteError(c, err)
	}
	if streamAdded {
		streamCloudlet.Remove(in.Cloudlet.Key, streamer)
	}
	return nil
}

func UpdateCloudletStream(ctx context.Context, rc *RegionContext, obj *edgeproto.Cloudlet, cb func(res *edgeproto.Result)) error {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.Key.Organization,
			ResourceCloudlets, ActionManage); err != nil {
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
	api := edgeproto.NewCloudletApiClient(rc.conn)
	stream, err := api.UpdateCloudlet(ctx, obj)
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

func UpdateCloudletObj(ctx context.Context, rc *RegionContext, obj *edgeproto.Cloudlet) ([]edgeproto.Result, error) {
	arr := []edgeproto.Result{}
	err := UpdateCloudletStream(ctx, rc, obj, func(res *edgeproto.Result) {
		arr = append(arr, *res)
	})
	return arr, err
}

func ShowCloudlet(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionCloudlet{}
	success, err := ReadConn(c, &in)
	if !success {
		return err
	}
	defer CloseConn(c)
	rc.region = in.Region

	err = ShowCloudletStream(ctx, rc, &in.Cloudlet, func(res *edgeproto.Cloudlet) {
		payload := ormapi.StreamPayload{}
		payload.Data = res
		WriteStream(c, &payload)
	})
	if err != nil {
		WriteError(c, err)
	}
	return nil
}

type ShowCloudletAuthz interface {
	Ok(obj *edgeproto.Cloudlet) bool
}

func ShowCloudletStream(ctx context.Context, rc *RegionContext, obj *edgeproto.Cloudlet, cb func(res *edgeproto.Cloudlet)) error {
	var authz ShowCloudletAuthz
	var err error
	if !rc.skipAuthz {
		authz, err = newShowCloudletAuthz(ctx, rc.region, rc.username, ResourceCloudlets, ActionView)
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
	api := edgeproto.NewCloudletApiClient(rc.conn)
	stream, err := api.ShowCloudlet(ctx, obj)
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
			if !authz.Ok(res) {
				continue
			}
		}
		cb(res)
	}
	return nil
}

func ShowCloudletObj(ctx context.Context, rc *RegionContext, obj *edgeproto.Cloudlet) ([]edgeproto.Cloudlet, error) {
	arr := []edgeproto.Cloudlet{}
	err := ShowCloudletStream(ctx, rc, obj, func(res *edgeproto.Cloudlet) {
		arr = append(arr, *res)
	})
	return arr, err
}

func GetCloudletManifest(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionCloudlet{}
	if err := c.Bind(&in); err != nil {
		return bindErr(c, err)
	}
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("org", in.Cloudlet.Key.Organization)
	resp, err := GetCloudletManifestObj(ctx, rc, &in.Cloudlet)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			err = fmt.Errorf("%s", st.Message())
		}
	}
	return setReply(c, err, resp)
}

func GetCloudletManifestObj(ctx context.Context, rc *RegionContext, obj *edgeproto.Cloudlet) (*edgeproto.CloudletManifest, error) {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.Key.Organization,
			ResourceCloudlets, ActionManage); err != nil {
			return nil, err
		}
	}
	if rc.conn == nil {
		conn, err := connectController(ctx, rc.region)
		if err != nil {
			return nil, err
		}
		rc.conn = conn
		defer func() {
			rc.conn.Close()
			rc.conn = nil
		}()
	}
	api := edgeproto.NewCloudletApiClient(rc.conn)
	return api.GetCloudletManifest(ctx, obj)
}

func GetCloudletProps(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionCloudletProps{}
	if err := c.Bind(&in); err != nil {
		return bindErr(c, err)
	}
	rc.region = in.Region
	resp, err := GetCloudletPropsObj(ctx, rc, &in.CloudletProps)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			err = fmt.Errorf("%s", st.Message())
		}
	}
	return setReply(c, err, resp)
}

func GetCloudletPropsObj(ctx context.Context, rc *RegionContext, obj *edgeproto.CloudletProps) (*edgeproto.CloudletProps, error) {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, "",
			ResourceCloudlets, ActionView); err != nil {
			return nil, err
		}
	}
	if rc.conn == nil {
		conn, err := connectController(ctx, rc.region)
		if err != nil {
			return nil, err
		}
		rc.conn = conn
		defer func() {
			rc.conn.Close()
			rc.conn = nil
		}()
	}
	api := edgeproto.NewCloudletApiClient(rc.conn)
	return api.GetCloudletProps(ctx, obj)
}

func AddCloudletResMapping(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionCloudletResMap{}
	if err := c.Bind(&in); err != nil {
		return bindErr(c, err)
	}
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("org", in.CloudletResMap.Key.Organization)
	resp, err := AddCloudletResMappingObj(ctx, rc, &in.CloudletResMap)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			err = fmt.Errorf("%s", st.Message())
		}
	}
	return setReply(c, err, resp)
}

func AddCloudletResMappingObj(ctx context.Context, rc *RegionContext, obj *edgeproto.CloudletResMap) (*edgeproto.Result, error) {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.Key.Organization,
			ResourceCloudlets, ActionManage); err != nil {
			return nil, err
		}
	}
	if rc.conn == nil {
		conn, err := connectController(ctx, rc.region)
		if err != nil {
			return nil, err
		}
		rc.conn = conn
		defer func() {
			rc.conn.Close()
			rc.conn = nil
		}()
	}
	api := edgeproto.NewCloudletApiClient(rc.conn)
	return api.AddCloudletResMapping(ctx, obj)
}

func RemoveCloudletResMapping(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionCloudletResMap{}
	if err := c.Bind(&in); err != nil {
		return bindErr(c, err)
	}
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("org", in.CloudletResMap.Key.Organization)
	resp, err := RemoveCloudletResMappingObj(ctx, rc, &in.CloudletResMap)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			err = fmt.Errorf("%s", st.Message())
		}
	}
	return setReply(c, err, resp)
}

func RemoveCloudletResMappingObj(ctx context.Context, rc *RegionContext, obj *edgeproto.CloudletResMap) (*edgeproto.Result, error) {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.Key.Organization,
			ResourceCloudlets, ActionManage); err != nil {
			return nil, err
		}
	}
	if rc.conn == nil {
		conn, err := connectController(ctx, rc.region)
		if err != nil {
			return nil, err
		}
		rc.conn = conn
		defer func() {
			rc.conn.Close()
			rc.conn = nil
		}()
	}
	api := edgeproto.NewCloudletApiClient(rc.conn)
	return api.RemoveCloudletResMapping(ctx, obj)
}

func FindFlavorMatch(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionFlavorMatch{}
	if err := c.Bind(&in); err != nil {
		return bindErr(c, err)
	}
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("org", in.FlavorMatch.Key.Organization)
	resp, err := FindFlavorMatchObj(ctx, rc, &in.FlavorMatch)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			err = fmt.Errorf("%s", st.Message())
		}
	}
	return setReply(c, err, resp)
}

func FindFlavorMatchObj(ctx context.Context, rc *RegionContext, obj *edgeproto.FlavorMatch) (*edgeproto.FlavorMatch, error) {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.Key.Organization,
			ResourceCloudlets, ActionView); err != nil {
			return nil, err
		}
	}
	if rc.conn == nil {
		conn, err := connectController(ctx, rc.region)
		if err != nil {
			return nil, err
		}
		rc.conn = conn
		defer func() {
			rc.conn.Close()
			rc.conn = nil
		}()
	}
	api := edgeproto.NewCloudletApiClient(rc.conn)
	return api.FindFlavorMatch(ctx, obj)
}

func ShowCloudletInfo(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionCloudletInfo{}
	success, err := ReadConn(c, &in)
	if !success {
		return err
	}
	defer CloseConn(c)
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("org", in.CloudletInfo.Key.Organization)

	err = ShowCloudletInfoStream(ctx, rc, &in.CloudletInfo, func(res *edgeproto.CloudletInfo) {
		payload := ormapi.StreamPayload{}
		payload.Data = res
		WriteStream(c, &payload)
	})
	if err != nil {
		WriteError(c, err)
	}
	return nil
}

func ShowCloudletInfoStream(ctx context.Context, rc *RegionContext, obj *edgeproto.CloudletInfo, cb func(res *edgeproto.CloudletInfo)) error {
	var authz *AuthzShow
	var err error
	if !rc.skipAuthz {
		authz, err = newShowAuthz(ctx, rc.region, rc.username, ResourceCloudletAnalytics, ActionView)
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
	api := edgeproto.NewCloudletInfoApiClient(rc.conn)
	stream, err := api.ShowCloudletInfo(ctx, obj)
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

func ShowCloudletInfoObj(ctx context.Context, rc *RegionContext, obj *edgeproto.CloudletInfo) ([]edgeproto.CloudletInfo, error) {
	arr := []edgeproto.CloudletInfo{}
	err := ShowCloudletInfoStream(ctx, rc, obj, func(res *edgeproto.CloudletInfo) {
		arr = append(arr, *res)
	})
	return arr, err
}

func InjectCloudletInfo(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionCloudletInfo{}
	if err := c.Bind(&in); err != nil {
		return bindErr(c, err)
	}
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("org", in.CloudletInfo.Key.Organization)
	resp, err := InjectCloudletInfoObj(ctx, rc, &in.CloudletInfo)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			err = fmt.Errorf("%s", st.Message())
		}
	}
	return setReply(c, err, resp)
}

func InjectCloudletInfoObj(ctx context.Context, rc *RegionContext, obj *edgeproto.CloudletInfo) (*edgeproto.Result, error) {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.Key.Organization,
			ResourceCloudlets, ActionManage); err != nil {
			return nil, err
		}
	}
	if rc.conn == nil {
		conn, err := connectController(ctx, rc.region)
		if err != nil {
			return nil, err
		}
		rc.conn = conn
		defer func() {
			rc.conn.Close()
			rc.conn = nil
		}()
	}
	api := edgeproto.NewCloudletInfoApiClient(rc.conn)
	return api.InjectCloudletInfo(ctx, obj)
}

func EvictCloudletInfo(c echo.Context) error {
	ctx := GetContext(c)
	rc := &RegionContext{}
	claims, err := getClaims(c)
	if err != nil {
		return err
	}
	rc.username = claims.Username

	in := ormapi.RegionCloudletInfo{}
	if err := c.Bind(&in); err != nil {
		return bindErr(c, err)
	}
	rc.region = in.Region
	span := log.SpanFromContext(ctx)
	span.SetTag("org", in.CloudletInfo.Key.Organization)
	resp, err := EvictCloudletInfoObj(ctx, rc, &in.CloudletInfo)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			err = fmt.Errorf("%s", st.Message())
		}
	}
	return setReply(c, err, resp)
}

func EvictCloudletInfoObj(ctx context.Context, rc *RegionContext, obj *edgeproto.CloudletInfo) (*edgeproto.Result, error) {
	log.SetContextTags(ctx, edgeproto.GetTags(obj))
	if !rc.skipAuthz {
		if err := authorized(ctx, rc.username, obj.Key.Organization,
			ResourceCloudlets, ActionManage); err != nil {
			return nil, err
		}
	}
	if rc.conn == nil {
		conn, err := connectController(ctx, rc.region)
		if err != nil {
			return nil, err
		}
		rc.conn = conn
		defer func() {
			rc.conn.Close()
			rc.conn = nil
		}()
	}
	api := edgeproto.NewCloudletInfoApiClient(rc.conn)
	return api.EvictCloudletInfo(ctx, obj)
}
