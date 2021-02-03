// Code generated by candi v1.3.1.

package grpchandler

import (
	"context"

	proto "monorepo/sdk/auth-service/proto/token"
	"monorepo/services/auth-service/internal/modules/token/domain"
	"monorepo/services/auth-service/internal/modules/token/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"pkg.agungdp.dev/candi/codebase/factory/types"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
)

// GRPCHandler rpc handler
type GRPCHandler struct {
	mw        interfaces.Middleware
	uc        usecase.TokenUsecase
	validator interfaces.Validator
}

// NewGRPCHandler func
func NewGRPCHandler(mw interfaces.Middleware, uc usecase.TokenUsecase, validator interfaces.Validator) *GRPCHandler {
	return &GRPCHandler{
		mw: mw, uc: uc, validator: validator,
	}
}

// Register grpc server
func (h *GRPCHandler) Register(server *grpc.Server, mwGroup *types.MiddlewareGroup) {
	proto.RegisterTokenHandlerServer(server, h)

	// register middleware for method
	mwGroup.AddProto(proto.File_token_token_proto, "GenerateToken", h.mw.GRPCBasicAuth)
	// mwGroup.AddProto(proto.File_token_token_proto, "ValidateToken", h.mw.GRPCBasicAuth)
}

// GenerateToken rpc
func (h *GRPCHandler) GenerateToken(ctx context.Context, req *proto.UserData) (*proto.ResponseGenerate, error) {
	trace := tracer.StartTrace(ctx, "TokenDeliveryGRPC:GenerateToken")
	defer trace.Finish()
	ctx = trace.Context()

	var tokenClaim domain.Claim
	tokenClaim.User.ID = req.ID
	tokenClaim.User.Username = req.Username

	result := <-h.uc.Generate(ctx, &tokenClaim)
	if result.Error != nil {
		return nil, grpc.Errorf(codes.Internal, "%v", result.Error)
	}

	tokenString := result.Data.(string)

	return &proto.ResponseGenerate{
		Success: true,
		Data: &proto.ResponseGenerate_Token{
			Token: tokenString,
		},
	}, nil
}

// ValidateToken rpc
func (h *GRPCHandler) ValidateToken(ctx context.Context, req *proto.PayloadValidate) (*proto.ResponseValidation, error) {
	trace := tracer.StartTrace(ctx, "TokenDeliveryGRPC:ValidateToken")
	defer trace.Finish()
	ctx = trace.Context()

	result := <-h.uc.Validate(ctx, req.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	claim := result.Data.(*domain.Claim)

	return &proto.ResponseValidation{
		Success: true,
		Claim: &proto.ResponseValidation_ClaimData{
			Audience:  claim.Audience,
			Subject:   claim.Subject,
			ExpiresAt: claim.ExpiresAt,
			User: &proto.UserData{
				ID:       claim.User.ID,
				Username: claim.User.Username,
			},
		},
	}, nil
}
