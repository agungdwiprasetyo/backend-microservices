// Code generated by candi v1.3.1.

package workerhandler

import (
	"context"
	"fmt"

	"monorepo/services/line-chatbot/internal/modules/chatbot/usecase"

	"pkg.agungdwiprasetyo.com/candi/codebase/factory/types"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
	"pkg.agungdwiprasetyo.com/candi/tracer"
)

// RedisHandler struct
type RedisHandler struct {
	uc        usecase.ChatbotUsecase
	validator interfaces.Validator
}

// NewRedisHandler constructor
func NewRedisHandler(uc usecase.ChatbotUsecase, validator interfaces.Validator) *RedisHandler {
	return &RedisHandler{
		uc:        uc,
		validator: validator,
	}
}

// MountHandlers mount handler group
func (h *RedisHandler) MountHandlers(group *types.WorkerHandlerGroup) {
	group.Add("chatbot-sample", h.handleChatbot)
}

func (h *RedisHandler) handleChatbot(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "ChatbotDeliveryRedis:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	fmt.Println("redis subs: execute sample")
	return nil
}
