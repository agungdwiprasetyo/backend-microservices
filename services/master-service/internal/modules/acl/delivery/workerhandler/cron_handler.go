// Code generated by candi v1.4.0.

package workerhandler

import (
	"context"
	"fmt"

	"monorepo/services/master-service/internal/modules/acl/usecase"

	"pkg.agungdp.dev/candi/candihelper"
	"pkg.agungdp.dev/candi/codebase/factory/types"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
)

// CronHandler struct
type CronHandler struct {
	uc        usecase.ACLUsecase
	validator interfaces.Validator
}

// NewCronHandler constructor
func NewCronHandler(uc usecase.ACLUsecase, validator interfaces.Validator) *CronHandler {
	return &CronHandler{
		uc:        uc,
		validator: validator,
	}
}

// MountHandlers mount handler group
func (h *CronHandler) MountHandlers(group *types.WorkerHandlerGroup) {
	group.Add(candihelper.CronJobKeyToString("acl-scheduler", "10s"), h.handleACL)
}

func (h *CronHandler) handleACL(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "AclDeliveryCron:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	fmt.Println("cron: execute in module acl")
	return nil
}
