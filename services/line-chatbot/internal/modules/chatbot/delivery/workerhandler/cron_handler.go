// Code generated by candi v1.3.1.

package workerhandler

import (
	"context"
	"fmt"

	"monorepo/services/line-chatbot/internal/modules/chatbot/usecase"

	"pkg.agungdwiprasetyo.com/candi/candihelper"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/types"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
	"pkg.agungdwiprasetyo.com/candi/tracer"
)

// CronHandler struct
type CronHandler struct {
	uc        usecase.ChatbotUsecase
	validator interfaces.Validator
}

// NewCronHandler constructor
func NewCronHandler(uc usecase.ChatbotUsecase, validator interfaces.Validator) *CronHandler {
	return &CronHandler{
		uc:        uc,
		validator: validator,
	}
}

// MountHandlers mount handler group
func (h *CronHandler) MountHandlers(group *types.WorkerHandlerGroup) {
	group.Add(candihelper.CronJobKeyToString("chatbot-scheduler", "10s"), h.handleChatbot)
}

func (h *CronHandler) handleChatbot(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "ChatbotDeliveryCron:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	fmt.Println("cron: execute in module chatbot")
	return nil
}
