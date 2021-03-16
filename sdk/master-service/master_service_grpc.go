package masterservice

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	aclproto "monorepo/sdk/master-service/proto/acl"
	appsproto "monorepo/sdk/master-service/proto/apps"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"pkg.agungdp.dev/candi/logger"
	"pkg.agungdp.dev/candi/tracer"
)

type masterServiceGRPCImpl struct {
	host       string
	authKey    string
	aclClient  aclproto.AclHandlerClient
	appsClient appsproto.AppsHandlerClient
}

// NewMasterServiceGRPC constructor
func NewMasterServiceGRPC(host string, authKey string) MasterService {

	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithConnectParams(grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay:  50 * time.Millisecond,
			Multiplier: 5,
			MaxDelay:   50 * time.Millisecond,
		},
		MinConnectTimeout: 1 * time.Second,
	}))
	if err != nil {
		panic(err)
	}

	return &masterServiceGRPCImpl{
		host:       host,
		authKey:    authKey,
		aclClient:  aclproto.NewAclHandlerClient(conn),
		appsClient: appsproto.NewAppsHandlerClient(conn),
	}
}

func (a *masterServiceGRPCImpl) CheckPermission(ctx context.Context, userID string, permissionCode string) (role string, err error) {
	trace := tracer.StartTrace(ctx, "MasterServiceSDK:CheckPermission")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
		trace.Finish()
	}()
	ctx = trace.Context()

	md := metadata.Pairs("authorization", a.authKey)
	trace.InjectGRPCMetadata(md)
	ctx = metadata.NewOutgoingContext(ctx, md)
	reqData := &aclproto.CheckPermissionRequest{
		UserID: userID, PermissionCode: permissionCode,
	}

	trace.SetTag("metadata", md)
	trace.SetTag("host", a.host)
	tracer.Log(ctx, "request.data", reqData)

	resp, err := a.aclClient.CheckPermission(ctx, reqData)
	if err != nil {
		trace.SetError(err)
		logger.LogE(err.Error())
		desc, ok := status.FromError(err)
		if ok {
			err = errors.New(desc.Message())
		}
		return
	}
	tracer.Log(ctx, "response.data", resp)

	role = resp.RoleID
	return
}

func (a *masterServiceGRPCImpl) GetUserApps(ctx context.Context, userID string) (userApps []UserApps, err error) {
	trace := tracer.StartTrace(ctx, "MasterServiceSDK:GetUserApps")
	defer func() { trace.SetError(err); trace.Finish() }()
	ctx = trace.Context()

	md := metadata.Pairs("authorization", a.authKey)
	trace.InjectGRPCMetadata(md)

	ctx = metadata.NewOutgoingContext(ctx, md)
	reqData := &appsproto.RequestUserApps{
		UserID: userID,
	}

	trace.SetTag("metadata", md)
	trace.SetTag("host", a.host)
	tracer.Log(ctx, "request.data", reqData)

	stream, err := a.appsClient.GetUserApps(ctx, reqData)
	if err != nil {
		logger.LogE(err.Error())
		return userApps, err
	}
	defer stream.CloseSend()

	// stream get data with grpc
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return userApps, err
		}

		userApp := UserApps{
			ID: msg.ID, Code: msg.Code, Name: msg.Name, Icon: msg.Icon, FrontendURL: msg.FrontendUrl, BackendURL: msg.BackendUrl,
		}
		userApp.Role.ID = msg.Role.ID
		userApp.Role.Code = msg.Role.Code
		userApp.Role.Name = msg.Role.Name
		userApps = append(userApps, userApp)
	}

	return
}
