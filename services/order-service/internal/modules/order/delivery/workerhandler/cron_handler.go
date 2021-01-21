// Code generated by candi v1.3.3.

package workerhandler

import (
	"context"
	"fmt"

	"monorepo/services/order-service/internal/modules/order/usecase"

	"pkg.agungdwiprasetyo.com/candi/candihelper"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/types"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
	"pkg.agungdwiprasetyo.com/candi/tracer"
)

// CronHandler struct
type CronHandler struct {
	uc        usecase.OrderUsecase
	validator interfaces.Validator
}

// NewCronHandler constructor
func NewCronHandler(uc usecase.OrderUsecase, validator interfaces.Validator) *CronHandler {
	return &CronHandler{
		uc:        uc,
		validator: validator,
	}
}

// MountHandlers mount handler group
func (h *CronHandler) MountHandlers(group *types.WorkerHandlerGroup) {
	group.Add(candihelper.CronJobKeyToString("order-scheduler", "10s"), h.handleOrder)
}

func (h *CronHandler) handleOrder(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "OrderDeliveryCron:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	fmt.Println("cron: execute in module order")
	h.uc.Hello(ctx)
	return nil
}
