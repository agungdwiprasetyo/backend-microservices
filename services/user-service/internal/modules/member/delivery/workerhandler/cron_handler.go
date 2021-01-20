// Code generated by candi v1.3.1.

package workerhandler

import (
	"context"
	"fmt"

	"monorepo/services/user-service/internal/modules/member/usecase"

	"pkg.agungdwiprasetyo.com/candi/candihelper"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/types"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
	"pkg.agungdwiprasetyo.com/candi/tracer"
)

// CronHandler struct
type CronHandler struct {
	uc        usecase.MemberUsecase
	validator interfaces.Validator
}

// NewCronHandler constructor
func NewCronHandler(uc usecase.MemberUsecase, validator interfaces.Validator) *CronHandler {
	return &CronHandler{
		uc:        uc,
		validator: validator,
	}
}

// MountHandlers mount handler group
func (h *CronHandler) MountHandlers(group *types.WorkerHandlerGroup) {
	group.Add(candihelper.CronJobKeyToString("member-scheduler", "10s"), h.handleMember)
}

func (h *CronHandler) handleMember(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "MemberDeliveryCron:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	fmt.Println("cron: execute in module member")
	h.uc.Hello(ctx)
	return nil
}