// Code generated by candi v1.3.1.

package workerhandler

import (
	"context"
	"fmt"

	"monorepo/services/auth-service/internal/modules/token/usecase"

	"pkg.agungdp.dev/candi/codebase/factory/types"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
)

// RedisHandler struct
type RedisHandler struct {
	uc        usecase.TokenUsecase
	validator interfaces.Validator
}

// NewRedisHandler constructor
func NewRedisHandler(uc usecase.TokenUsecase, validator interfaces.Validator) *RedisHandler {
	return &RedisHandler{
		uc:        uc,
		validator: validator,
	}
}

// MountHandlers mount handler group
func (h *RedisHandler) MountHandlers(group *types.WorkerHandlerGroup) {
	group.Add("token-sample", h.handleToken)
}

func (h *RedisHandler) handleToken(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "TokenDeliveryRedis:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	fmt.Println("redis subs: execute sample")
	// h.uc.Hello(ctx)
	return nil
}
