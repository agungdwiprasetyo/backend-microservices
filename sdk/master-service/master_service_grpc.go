package masterservice

import (
	"context"
	"errors"
	"fmt"
	"time"

	proto "monorepo/sdk/master-service/proto/acl"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"pkg.agungdp.dev/candi/logger"
	"pkg.agungdp.dev/candi/tracer"
)

type masterServiceGRPCImpl struct {
	host    string
	authKey string
	client  proto.AclHandlerClient
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
		host:    host,
		authKey: authKey,
		client:  proto.NewAclHandlerClient(conn),
	}
}

func (a *masterServiceGRPCImpl) CheckPermission(ctx context.Context, userID string, permissionCode string) (err error) {
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
	reqData := &proto.CheckPermissionRequest{
		UserID: userID, PermissionCode: permissionCode,
	}

	trace.SetTag("metadata", md)
	trace.SetTag("host", a.host)
	tracer.Log(ctx, "request.data", reqData)

	resp, err := a.client.CheckPermission(ctx, reqData)
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

	if !resp.IsAllowed {
		return errors.New("Not allowed")
	}

	return nil
}
