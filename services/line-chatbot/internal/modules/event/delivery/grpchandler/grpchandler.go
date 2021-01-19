// Code generated by candi v1.3.1.

package grpchandler

import (
	"context"

	proto "monorepo/sdk/line-chatbot/proto/event"
	"monorepo/services/line-chatbot/internal/modules/event/usecase"

	"google.golang.org/grpc"

	"pkg.agungdwiprasetyo.com/candi/candishared"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/types"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
	"pkg.agungdwiprasetyo.com/candi/tracer"
)

// GRPCHandler rpc handler
type GRPCHandler struct {
	mw        interfaces.Middleware
	uc        usecase.EventUsecase
	validator interfaces.Validator
}

// NewGRPCHandler func
func NewGRPCHandler(mw interfaces.Middleware, uc usecase.EventUsecase, validator interfaces.Validator) *GRPCHandler {
	return &GRPCHandler{
		mw: mw, uc: uc, validator: validator,
	}
}

// Register grpc server
func (h *GRPCHandler) Register(server *grpc.Server, mwGroup *types.MiddlewareGroup) {
	proto.RegisterEventHandlerServer(server, h)

	// register middleware for method
	mwGroup.AddProto(proto.File_event_event_proto, "Hello", h.mw.GRPCBearerAuth)
}

// Hello rpc method
func (h *GRPCHandler) Hello(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	trace := tracer.StartTrace(ctx, "EventDeliveryGRPC:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using GRPCBearerAuth in middleware for this handler

	return &proto.Response{
		Message: "hello, with your session (" + tokenClaim.Audience + ")",
	}, nil
}
