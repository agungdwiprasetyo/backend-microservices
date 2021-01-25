// Code generated by candi v1.3.1.

package usecase

import (
	"context"
	"fmt"
	"strings"

	eventdomain "monorepo/services/line-chatbot/internal/modules/event/domain"
	"monorepo/services/line-chatbot/pkg/helper"
	shareddomain "monorepo/services/line-chatbot/pkg/shared/domain"
	linebotapi "monorepo/services/line-chatbot/pkg/shared/linebot-api"
	"monorepo/services/line-chatbot/pkg/shared/repository"
	"monorepo/services/line-chatbot/pkg/shared/translator"

	"github.com/line/line-bot-sdk-go/linebot"
	"pkg.agungdwiprasetyo.com/candi/candihelper"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/dependency"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
	"pkg.agungdwiprasetyo.com/candi/tracer"
)

type chatbotUsecaseImpl struct {
	cache      interfaces.Cache
	repoMongo  *repository.RepoMongo
	translator translator.Translator
	botAPI     linebotapi.Linebot
	lineClient *linebot.Client
}

// NewChatbotUsecase usecase impl constructor
func NewChatbotUsecase(deps dependency.Dependency) ChatbotUsecase {
	return &chatbotUsecaseImpl{
		cache:      deps.GetRedisPool().Cache(),
		repoMongo:  repository.GetSharedRepoMongo(),
		botAPI:     linebotapi.NewLineBotHTTP(),
		translator: translator.NewTranslatorHTTP(),
		lineClient: deps.GetExtended()[helper.LineClient].(*linebot.Client),
	}
}

func (uc *chatbotUsecaseImpl) ProcessCallback(ctx context.Context, events []*linebot.Event) error {
	trace := tracer.StartTrace(ctx, "ChatbotUsecase:ProcessCallback")
	defer func() {
		if r := recover(); r != nil {
			trace.SetError(fmt.Errorf("%v", r))
		}
		trace.Finish()
	}()
	ctx = trace.Context()

	for _, event := range events {
		var (
			profile  eventdomain.Profile
			eventLog eventdomain.Event
		)

		profileResp, err := uc.lineClient.GetProfile(event.Source.UserID).Do()
		if err == nil {
			profile.Type = string(event.Source.Type)
			switch profile.Type {
			case "user":
				profile.ID = event.Source.UserID
			case "group":
				profile.ID = event.Source.GroupID
			}
			profile.Name = profileResp.DisplayName
			profile.Avatar = profileResp.PictureURL
			profile.StatusMessage = profileResp.StatusMessage
		}

		// save event
		eventLog.ReplyToken = event.ReplyToken
		eventLog.Type = string(event.Type)
		eventLog.Timestamp = event.Timestamp
		eventLog.SourceID = profile.ID
		eventLog.SourceType = profile.Type

		switch event.Type {
		case linebot.EventTypeJoin:
			uc.ReplyMessage(ctx, event, fmt.Sprintf("Hello %s :)", profileResp.DisplayName))

		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:

				var responseText string
				var inputText = message.Text
				translateToEnglish, translateToIndonesian := "terjemahkan ini ke inggris:", "terjemahkan ini ke indonesia:"

				switch {
				case strings.HasPrefix(strings.ToLower(inputText), translateToEnglish):
					i := strings.Index(strings.ToLower(inputText), translateToEnglish)
					if i >= 0 {
						inputText = inputText[i+len(translateToEnglish):]
					}
					responseText = uc.translator.Translate(ctx, "id", "en", inputText)

				case strings.HasPrefix(strings.ToLower(inputText), translateToIndonesian):
					i := strings.Index(strings.ToLower(inputText), translateToIndonesian)
					if i >= 0 {
						inputText = inputText[i+len(translateToIndonesian):]
					}
					responseText = uc.translator.Translate(ctx, "en", "id", inputText)

				default:
					responseText = uc.botAPI.ProcessText(ctx, inputText)
				}

				responseText = strings.TrimSpace(responseText)
				err := uc.ReplyMessage(ctx, event, responseText)
				if err != nil {
					eventLog.Error = candihelper.ToStringPtr(err.Error())
				}

				eventLog.Message.ID = message.ID
				eventLog.Message.Text = message.Text
				eventLog.Message.Response = responseText
			}
		}

		<-uc.repoMongo.EventRepo.Save(ctx, &eventLog)
		// uc.repo.Profile.Save(ctx, &profile)
	}

	return nil
}

func (uc *chatbotUsecaseImpl) ReplyMessage(ctx context.Context, event *linebot.Event, messages ...string) error {
	trace := tracer.StartTrace(ctx, "ChatbotUsecase:ReplyMessage")
	defer trace.Finish()
	ctx = trace.Context()

	tracer.Log(ctx, "event", event)
	tracer.Log(ctx, "messages", messages)

	var lineMessages []linebot.SendingMessage
	for _, msg := range messages {
		lineMessages = append(lineMessages, linebot.NewTextMessage(msg))
	}

	if _, err := uc.lineClient.ReplyMessage(event.ReplyToken, lineMessages...).Do(); err != nil {
		return err
	}

	return nil
}

func (uc *chatbotUsecaseImpl) PushMessageToChannel(ctx context.Context, to, title, message string) error {
	var lineMessage shareddomain.LineMessage

	lineMessage.To = to
	lineMessage.Messages = append(lineMessage.Messages, shareddomain.LineContentMessage{
		Type: "flex", AltText: title, Contents: shareddomain.LineContentFormat{
			Type: "bubble", Body: shareddomain.LineContentBody{
				Type: "box", Layout: "horizontal", Contents: []shareddomain.LineContent{
					{Type: "text", Text: message},
				},
			},
		},
	})

	return uc.botAPI.PushMessage(ctx, &lineMessage)
}
