// Code generated by candi v1.3.1.

package usecase

import (
	"sync"
	authusecase "monorepo/services/user-service/internal/modules/auth/usecase"
	memberusecase "monorepo/services/user-service/internal/modules/member/usecase"

	"pkg.agungdwiprasetyo.com/candi/codebase/factory/dependency"
)

type (
	// Usecase unit of work for all usecase in modules
	Usecase interface {
		Auth() authusecase.AuthUsecase
		Member() memberusecase.MemberUsecase
	}

	usecaseUow struct {
		auth authusecase.AuthUsecase
		member memberusecase.MemberUsecase
	}
)

var usecaseInst *usecaseUow
var once sync.Once

// SetSharedUsecase set singleton usecase unit of work instance
func SetSharedUsecase(deps dependency.Dependency) {
	once.Do(func() {
		usecaseInst = &usecaseUow{
			auth: authusecase.NewAuthUsecase(deps),
			member: memberusecase.NewMemberUsecase(deps),
		}
	})
}

// GetSharedUsecase get usecase unit of work instance
func GetSharedUsecase() Usecase {
	return usecaseInst
}
func (uc *usecaseUow) Auth() authusecase.AuthUsecase {
	return uc.auth
}
func (uc *usecaseUow) Member() memberusecase.MemberUsecase {
	return uc.member
}