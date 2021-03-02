package common

import (
	"context"
	"monorepo/services/master-service/internal/modules/apps/domain"
)

var commonUC Usecase

// Usecase common abstraction for bridging shared method inter usecase in module
type Usecase interface {
	// method from another usecase
	GetDetailApp(ctx context.Context, appID string) (data domain.AppDetail, err error)
}

// SetCommonUsecase constructor
func SetCommonUsecase(uc Usecase) {
	commonUC = uc
}

// GetCommonUsecase get common usecase
func GetCommonUsecase() Usecase {
	return commonUC
}
