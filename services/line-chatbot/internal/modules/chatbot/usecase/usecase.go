// Code generated by candi v1.3.1.

package usecase

import (
	"context"

	"github.com/line/line-bot-sdk-go/linebot"
)

// ChatbotUsecase abstraction
type ChatbotUsecase interface {
	ProcessCallback(ctx context.Context, events []*linebot.Event) error
	ReplyMessage(event *linebot.Event, messages ...string) error
	PushMessageToChannel(ctx context.Context, to, title, message string) error
}