// Code generated by candi v1.3.1.

package workerhandler

import (
	"context"
	"time"

	"monorepo/services/line-chatbot/internal/modules/chatbot/usecase"

	taskqueueworker "pkg.agungdp.dev/candi/codebase/app/task_queue_worker"
	"pkg.agungdp.dev/candi/codebase/factory/types"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
)

// TaskQueueHandler struct
type TaskQueueHandler struct {
	uc        usecase.ChatbotUsecase
	validator interfaces.Validator
}

// NewTaskQueueHandler constructor
func NewTaskQueueHandler(uc usecase.ChatbotUsecase, validator interfaces.Validator) *TaskQueueHandler {
	return &TaskQueueHandler{
		uc:        uc,
		validator: validator,
	}
}

// MountHandlers mount handler group
func (h *TaskQueueHandler) MountHandlers(group *types.WorkerHandlerGroup) {
	group.Add("chatbot-task-one", h.taskOne)
	group.Add("chatbot-task-two", h.taskTwo)
}

func (h *TaskQueueHandler) taskOne(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "ChatbotDeliveryTaskQueue:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	return &taskqueueworker.ErrorRetrier{
		Delay:   10 * time.Second,
		Message: "Error one",
	}
}

func (h *TaskQueueHandler) taskTwo(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "ChatbotDeliveryTaskQueue:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	return &taskqueueworker.ErrorRetrier{
		Delay:   3 * time.Second,
		Message: "Error two",
	}
}
