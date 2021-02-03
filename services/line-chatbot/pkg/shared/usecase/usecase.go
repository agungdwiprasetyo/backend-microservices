// Code generated by candi v1.3.1.

package usecase

import (
	chatbotusecase "monorepo/services/line-chatbot/internal/modules/chatbot/usecase"
	eventusecase "monorepo/services/line-chatbot/internal/modules/event/usecase"
	"sync"

	"pkg.agungdp.dev/candi/codebase/factory/dependency"
)

type (
	// Usecase unit of work for all usecase in modules
	Usecase interface {
		Chatbot() chatbotusecase.ChatbotUsecase
		Event() eventusecase.EventUsecase
	}

	usecaseUow struct {
		chatbot chatbotusecase.ChatbotUsecase
		event   eventusecase.EventUsecase
	}
)

var usecaseInst *usecaseUow
var once sync.Once

// SetSharedUsecase set singleton usecase unit of work instance
func SetSharedUsecase(deps dependency.Dependency) {
	once.Do(func() {
		usecaseInst = &usecaseUow{
			chatbot: chatbotusecase.NewChatbotUsecase(deps),
			event:   eventusecase.NewEventUsecase(deps),
		}
	})
}

// GetSharedUsecase get usecase unit of work instance
func GetSharedUsecase() Usecase {
	return usecaseInst
}
func (uc *usecaseUow) Chatbot() chatbotusecase.ChatbotUsecase {
	return uc.chatbot
}
func (uc *usecaseUow) Event() eventusecase.EventUsecase {
	return uc.event
}
