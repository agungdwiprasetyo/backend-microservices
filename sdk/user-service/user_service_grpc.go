package userservice

import (
	"context"
	"errors"
	"fmt"
	"time"

	proto "monorepo/sdk/user-service/proto/member"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"pkg.agungdp.dev/candi/logger"
	"pkg.agungdp.dev/candi/tracer"
)

type userServiceGRPCImpl struct {
	host    string
	authKey string
	client  proto.MemberHandlerClient
}

// NewUserServiceGRPC constructor
func NewUserServiceGRPC(host string, authKey string) UserService {

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

	return &userServiceGRPCImpl{
		host:    host,
		authKey: authKey,
		client:  proto.NewMemberHandlerClient(conn),
	}
}

func (a *userServiceGRPCImpl) GetMember(ctx context.Context, id string) (data Member, err error) {
	trace := tracer.StartTrace(ctx, "UserServiceSDK:GetMember")
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
	reqData := &proto.GetMemberRequest{
		ID: id,
	}

	trace.SetTag("metadata", md)
	trace.SetTag("host", a.host)
	tracer.Log(ctx, "request.data", reqData)

	resp, err := a.client.GetMember(ctx, reqData)
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

	data.ID = resp.ID
	data.Username = resp.Username
	data.Password = resp.Password
	data.Fullname = resp.Fullname
	data.CreatedAt = resp.CreatedAt
	data.ModifiedAt = resp.ModifiedAt
	return
}
