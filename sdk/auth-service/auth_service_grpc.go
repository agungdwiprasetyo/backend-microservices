package authservice

import (
	"context"
	"errors"
	"fmt"
	"time"

	proto "monorepo/sdk/auth-service/proto/token"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"pkg.agungdp.dev/candi/candishared"
	"pkg.agungdp.dev/candi/logger"
	"pkg.agungdp.dev/candi/tracer"
)

type authServiceGRPCImpl struct {
	host    string
	authKey string
	client  proto.TokenHandlerClient
}

// NewAuthServiceGRPC constructor
func NewAuthServiceGRPC(host string, authKey string) AuthService {

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

	return &authServiceGRPCImpl{
		host:    host,
		authKey: authKey,
		client:  proto.NewTokenHandlerClient(conn),
	}
}

func (a *authServiceGRPCImpl) ValidateToken(ctx context.Context, token string) (claim *candishared.TokenClaim, err error) {
	trace := tracer.StartTrace(ctx, "AuthServiceSDK:ValidateToken")
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
	reqData := &proto.PayloadValidate{
		Token: token,
	}

	trace.SetTag("metadata", md)
	trace.SetTag("host", a.host)
	tracer.Log(ctx, "request.data", reqData)

	resp, err := a.client.ValidateToken(ctx, reqData)
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

	claim = new(candishared.TokenClaim)
	claim.StandardClaims = jwt.StandardClaims{
		Audience:  resp.Claim.Audience,
		Subject:   resp.Claim.Subject,
		ExpiresAt: resp.Claim.ExpiresAt,
		Issuer:    resp.Claim.Issuer,
		IssuedAt:  resp.Claim.IssuedAt,
		NotBefore: resp.Claim.NotBefore,
	}
	claim.Additional = map[string]interface{}{
		"username": resp.Claim.User.Username,
		"user_id":  resp.Claim.User.ID,
	}

	return
}

func (a *authServiceGRPCImpl) GenerateToken(ctx context.Context, req PayloadGenerateToken) (token string, err error) {
	trace := tracer.StartTrace(ctx, "AuthServiceSDK:GenerateToken")
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
	reqData := &proto.UserData{
		ID: req.UserID, Username: req.Username,
	}

	trace.SetTag("metadata", md)
	trace.SetTag("host", a.host)
	tracer.Log(ctx, "request.data", reqData)

	resp, err := a.client.GenerateToken(ctx, reqData)
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

	token = resp.Data.Token
	return
}
