// Code generated by candi v1.3.1.

package usecase

import (
	"sync"
	pushnotifusecase "monorepo/services/notification-service/internal/modules/push-notif/usecase"

	"pkg.agungdwiprasetyo.com/candi/codebase/factory/dependency"
)

type (
	// Usecase unit of work for all usecase in modules
	Usecase interface {
		PushNotif() pushnotifusecase.PushNotifUsecase
	}

	usecaseUow struct {
		pushnotif pushnotifusecase.PushNotifUsecase
	}
)

var usecaseInst *usecaseUow
var once sync.Once

// SetSharedUsecase set singleton usecase unit of work instance
func SetSharedUsecase(deps dependency.Dependency) {
	once.Do(func() {
		usecaseInst = &usecaseUow{
			pushnotif: pushnotifusecase.NewPushNotifUsecase(deps),
		}
	})
}

// GetSharedUsecase get usecase unit of work instance
func GetSharedUsecase() Usecase {
	return usecaseInst
}
func (uc *usecaseUow) PushNotif() pushnotifusecase.PushNotifUsecase {
	return uc.pushnotif
}
