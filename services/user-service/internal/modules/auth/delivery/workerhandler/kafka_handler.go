// Code generated by candi v1.3.1.

package workerhandler

import (
	"context"
	"fmt"

	"monorepo/services/user-service/internal/modules/auth/usecase"

	"pkg.agungdp.dev/candi/codebase/factory/types"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
)

// KafkaHandler struct
type KafkaHandler struct {
	uc        usecase.AuthUsecase
	validator interfaces.Validator
}

// NewKafkaHandler constructor
func NewKafkaHandler(uc usecase.AuthUsecase, validator interfaces.Validator) *KafkaHandler {
	return &KafkaHandler{
		uc:        uc,
		validator: validator,
	}
}

// MountHandlers mount handler group
func (h *KafkaHandler) MountHandlers(group *types.WorkerHandlerGroup) {
	group.Add("auth", h.handleAuth) // handling topic "auth"
}

// ProcessMessage from kafka consumer
func (h *KafkaHandler) handleAuth(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "AuthDeliveryKafka:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	fmt.Printf("message consumed by module auth. message: %s\n", string(message))
	h.uc.Hello(ctx)
	return nil
}
