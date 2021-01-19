// Code generated by candi v1.3.1.

package grpchandler

import (
	"context"

	proto "monorepo/sdk/user-service/proto/auth"
	"monorepo/services/user-service/internal/modules/auth/usecase"

	"google.golang.org/grpc"

	"pkg.agungdwiprasetyo.com/candi/candishared"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/types"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
	"pkg.agungdwiprasetyo.com/candi/tracer"
)

// GRPCHandler rpc handler
type GRPCHandler struct {
	mw        interfaces.Middleware
	uc        usecase.AuthUsecase
	validator interfaces.Validator
}

// NewGRPCHandler func
func NewGRPCHandler(mw interfaces.Middleware, uc usecase.AuthUsecase, validator interfaces.Validator) *GRPCHandler {
	return &GRPCHandler{
		mw: mw, uc: uc, validator: validator,
	}
}

// Register grpc server
func (h *GRPCHandler) Register(server *grpc.Server, mwGroup *types.MiddlewareGroup) {
	proto.RegisterAuthHandlerServer(server, h)

	// register middleware for method
	mwGroup.AddProto(proto.File_auth_auth_proto, "Hello", h.mw.GRPCBearerAuth)
}

// Hello rpc method
func (h *GRPCHandler) Hello(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	trace := tracer.StartTrace(ctx, "AuthDeliveryGRPC:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using GRPCBearerAuth in middleware for this handler

	return &proto.Response{
		Message: h.uc.Hello(ctx) + ", with your session (" + tokenClaim.Audience + ")",
	}, nil
}
