// Code generated by candi v1.3.1.

package workerhandler

import (
	"context"
	"fmt"

	"monorepo/services/storage-service/internal/modules/storage/usecase"

	"pkg.agungdwiprasetyo.com/candi/codebase/factory/types"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
	"pkg.agungdwiprasetyo.com/candi/tracer"
)

// KafkaHandler struct
type KafkaHandler struct {
	uc        usecase.StorageUsecase
	validator interfaces.Validator
}

// NewKafkaHandler constructor
func NewKafkaHandler(uc usecase.StorageUsecase, validator interfaces.Validator) *KafkaHandler {
	return &KafkaHandler{
		uc:        uc,
		validator: validator,
	}
}

// MountHandlers mount handler group
func (h *KafkaHandler) MountHandlers(group *types.WorkerHandlerGroup) {
	group.Add("storage", h.handleStorage) // handling topic "storage"
}

// ProcessMessage from kafka consumer
func (h *KafkaHandler) handleStorage(ctx context.Context, message []byte) error {
	trace := tracer.StartTrace(ctx, "StorageDeliveryKafka:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	fmt.Printf("message consumed by module storage. message: %s\n", string(message))
	h.uc.Hello(ctx)
	return nil
}