// Code generated by candi v1.4.0.

package grpchandler

import (
	"context"

	proto "monorepo/sdk/master-service/proto/apps"
	"monorepo/services/master-service/internal/modules/apps/usecase"

	"google.golang.org/grpc"

	"pkg.agungdp.dev/candi/candishared"
	"pkg.agungdp.dev/candi/codebase/factory/types"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
)

// GRPCHandler rpc handler
type GRPCHandler struct {
	mw        interfaces.Middleware
	uc        usecase.AppsUsecase
	validator interfaces.Validator
}

// NewGRPCHandler func
func NewGRPCHandler(mw interfaces.Middleware, uc usecase.AppsUsecase, validator interfaces.Validator) *GRPCHandler {
	return &GRPCHandler{
		mw: mw, uc: uc, validator: validator,
	}
}

// Register grpc server
func (h *GRPCHandler) Register(server *grpc.Server, mwGroup *types.MiddlewareGroup) {
	proto.RegisterAppsHandlerServer(server, h)

	// register middleware for method
	mwGroup.AddProto(proto.File_apps_apps_proto, "Hello", h.mw.GRPCBearerAuth)
}

// Hello rpc method
func (h *GRPCHandler) Hello(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	trace := tracer.StartTrace(ctx, "AppsDeliveryGRPC:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using GRPCBearerAuth in middleware for this handler

	return &proto.Response{
		Message: ", with your session (" + tokenClaim.Audience + ")",
	}, nil
}