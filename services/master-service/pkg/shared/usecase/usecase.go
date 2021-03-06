// Code generated by candi v1.5.6.

package usecase

import (
	aclusecase "monorepo/services/master-service/internal/modules/acl/usecase"
	appsusecase "monorepo/services/master-service/internal/modules/apps/usecase"
	"monorepo/services/master-service/pkg/shared/usecase/common"
	"sync"

	"pkg.agungdp.dev/candi/codebase/factory/dependency"
)

type (
	// Usecase unit of work for all usecase in modules
	Usecase interface {
		ACL() aclusecase.ACLUsecase
		Apps() appsusecase.AppsUsecase
	}

	usecaseUow struct {
		aclusecase.ACLUsecase
		appsusecase.AppsUsecase
	}
)

var usecaseInst *usecaseUow
var once sync.Once

// SetSharedUsecase set singleton usecase unit of work instance
func SetSharedUsecase(deps dependency.Dependency) {
	once.Do(func() {
		usecaseInst = &usecaseUow{
			ACLUsecase:  aclusecase.NewACLUsecase(deps),
			AppsUsecase: appsusecase.NewAppsUsecase(deps),
		}
		common.SetCommonUsecase(usecaseInst)
	})
}

// GetSharedUsecase get usecase unit of work instance
func GetSharedUsecase() Usecase {
	return usecaseInst
}
func (uc *usecaseUow) ACL() aclusecase.ACLUsecase {
	return uc.ACLUsecase
}
func (uc *usecaseUow) Apps() appsusecase.AppsUsecase {
	return uc.AppsUsecase
}
