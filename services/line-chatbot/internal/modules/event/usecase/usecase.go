// Code generated by candi v1.3.1.

package usecase

import (
	"context"
	"monorepo/services/line-chatbot/internal/modules/event/domain"

	"pkg.agungdwiprasetyo.com/candi/candishared"
)

// EventUsecase abstraction
type EventUsecase interface {
	FindAll(ctx context.Context, filter *candishared.Filter) (events []domain.Event, meta candishared.Meta, err error)
}
