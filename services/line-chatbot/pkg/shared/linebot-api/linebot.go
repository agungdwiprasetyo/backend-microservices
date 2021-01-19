package linebotapi

import (
	"context"
	"monorepo/services/line-chatbot/pkg/shared/domain"
)

// Linebot abstract interface
type Linebot interface {
	ProcessText(ctx context.Context, text string) string
	PushMessage(ctx context.Context, message *domain.LineMessage) error
}
