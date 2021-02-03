// Code generated by candi v1.3.3.

package usecase

import (
	masterusecase "monorepo/services/order-service/internal/modules/master/usecase"
	orderusecase "monorepo/services/order-service/internal/modules/order/usecase"
	"sync"

	"pkg.agungdp.dev/candi/codebase/factory/dependency"
)

type (
	// Usecase unit of work for all usecase in modules
	Usecase interface {
		Master() masterusecase.MasterUsecase
		Order() orderusecase.OrderUsecase
	}

	usecaseUow struct {
		master masterusecase.MasterUsecase
		order  orderusecase.OrderUsecase
	}
)

var usecaseInst *usecaseUow
var once sync.Once

// SetSharedUsecase set singleton usecase unit of work instance
func SetSharedUsecase(deps dependency.Dependency) {
	once.Do(func() {
		usecaseInst = &usecaseUow{
			master: masterusecase.NewMasterUsecase(deps),
			order:  orderusecase.NewOrderUsecase(deps),
		}
	})
}

// GetSharedUsecase get usecase unit of work instance
func GetSharedUsecase() Usecase {
	return usecaseInst
}
func (uc *usecaseUow) Master() masterusecase.MasterUsecase {
	return uc.master
}
func (uc *usecaseUow) Order() orderusecase.OrderUsecase {
	return uc.order
}
