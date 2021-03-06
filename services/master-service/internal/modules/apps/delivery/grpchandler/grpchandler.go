// Code generated by candi v1.8.18.

package grpchandler

import (
	proto "monorepo/sdk/master-service/proto/apps"
	"monorepo/services/master-service/internal/modules/apps/usecase"

	"google.golang.org/grpc"

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
	mwGroup.AddProto(proto.File_apps_apps_proto, h.GetUserApps, h.mw.GRPCBasicAuth)
}

// GetUserApps rpc method
func (h *GRPCHandler) GetUserApps(req *proto.RequestUserApps, stream proto.AppsHandler_GetUserAppsServer) (err error) {
	trace := tracer.StartTrace(stream.Context(), "AppsDeliveryGRPC:GetUserApps")
	defer trace.Finish()
	ctx := trace.Context()

	userApps, err := h.uc.GetUserApps(ctx, req.UserID)
	if err != nil {
		return err
	}

	for _, userApp := range userApps {
		resp := &proto.ResponseUserApps{
			ID: userApp.ID, Code: userApp.Code, Name: userApp.Name, Icon: userApp.Icon, FrontendUrl: userApp.FrontendURL, BackendUrl: userApp.BackendURL,
			Role: &proto.ResponseUserApps_RoleType{
				Code: userApp.Role.Code,
				ID:   userApp.Role.ID,
				Name: userApp.Role.Name,
			},
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}

	return nil
}
