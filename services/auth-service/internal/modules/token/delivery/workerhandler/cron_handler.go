// Code generated by candi v1.3.1.

package workerhandler

import (
	"context"
	"fmt"

	"monorepo/services/auth-service/internal/modules/token/usecase"

	"pkg.agungdp.dev/candi/candihelper"
	"pkg.agungdp.dev/candi/codebase/factory/types"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
)

// CronHandler struct
type CronHandler struct {
	uc        usecase.TokenUsecase
	validator interfaces.Validator
}

// NewCronHandler constructor
func NewCronHandler(uc usecase.TokenUsecase, validator interfaces.Validator) *CronHandler {
	return &CronHandler{
		uc:        uc,
		validator: validator,
	}
}

// MountHandlers mount handler group
func (h *CronHandler) MountHandlers(group *types.WorkerHandlerGroup) {
	group.Add(candihelper.CronJobKeyToString("token-scheduler", "20:41:00@3s"), h.handleToken)
}

func (h *CronHandler) handleToken(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "TokenDeliveryCron:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	fmt.Println("cron: execute in module token")
	// h.uc.Hello(ctx)
	return nil
}
