// Code generated by candi v1.4.0.

package workerhandler

import (
	"context"
	"time"

	"monorepo/services/master-service/internal/modules/acl/usecase"

	taskqueueworker "pkg.agungdp.dev/candi/codebase/app/task_queue_worker"
	"pkg.agungdp.dev/candi/codebase/factory/types"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
)

// TaskQueueHandler struct
type TaskQueueHandler struct {
	uc        usecase.ACLUsecase
	validator interfaces.Validator
}

// NewTaskQueueHandler constructor
func NewTaskQueueHandler(uc usecase.ACLUsecase, validator interfaces.Validator) *TaskQueueHandler {
	return &TaskQueueHandler{
		uc:        uc,
		validator: validator,
	}
}

// MountHandlers mount handler group
func (h *TaskQueueHandler) MountHandlers(group *types.WorkerHandlerGroup) {
	group.Add("acl-task-one", h.taskOne)
	group.Add("acl-task-two", h.taskTwo)
}

func (h *TaskQueueHandler) taskOne(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "AclDeliveryTaskQueue:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	h.uc.Hello(ctx)

	return &taskqueueworker.ErrorRetrier{
		Delay:   10 * time.Second,
		Message: "Error one",
	}
}

func (h *TaskQueueHandler) taskTwo(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "AclDeliveryTaskQueue:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	return &taskqueueworker.ErrorRetrier{
		Delay:   3 * time.Second,
		Message: "Error two",
	}
}
