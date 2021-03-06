// Code generated by candi v1.3.1.

package usecase

import (
	"context"
	"monorepo/services/line-chatbot/internal/modules/chatbot/domain"

	"github.com/line/line-bot-sdk-go/linebot"
)

// ChatbotUsecase abstraction
type ChatbotUsecase interface {
	ProcessCallback(ctx context.Context, events []*linebot.Event) error
	ReplyMessage(ctx context.Context, event *linebot.Event, messages ...string) error
	PushMessageToChannel(ctx context.Context, payload domain.PushMessagePayload) error
}
