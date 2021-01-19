// Code generated by candi v1.3.1.

package workerhandler

import (
	"context"
	"fmt"

	"monorepo/services/notification-service/internal/modules/push-notif/usecase"

	"pkg.agungdwiprasetyo.com/candi/candihelper"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/types"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
	"pkg.agungdwiprasetyo.com/candi/tracer"
)

// CronHandler struct
type CronHandler struct {
	uc        usecase.PushNotifUsecase
	validator interfaces.Validator
}

// NewCronHandler constructor
func NewCronHandler(uc usecase.PushNotifUsecase, validator interfaces.Validator) *CronHandler {
	return &CronHandler{
		uc:        uc,
		validator: validator,
	}
}

// MountHandlers mount handler group
func (h *CronHandler) MountHandlers(group *types.WorkerHandlerGroup) {
	group.Add(candihelper.CronJobKeyToString("push-notif-scheduler", "10s"), h.handlePushNotif)
}

func (h *CronHandler) handlePushNotif(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "PushNotifDeliveryCron:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	fmt.Println("cron: execute in module push-notif")
	h.uc.Hello(ctx)
	return nil
}
