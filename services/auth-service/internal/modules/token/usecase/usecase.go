// Code generated by candi v1.8.17.

package usecase

import (
	"context"
	"monorepo/services/auth-service/internal/modules/token/domain"
)

// TokenUsecase abstraction
type TokenUsecase interface {
	Generate(ctx context.Context, payload *domain.Claim) (resp domain.ResponseGenerateToken, err error)
	Refresh(ctx context.Context, token, refreshToken string) (resp domain.ResponseGenerateToken, err error)
	Validate(ctx context.Context, tokenString string) (claim *domain.Claim, err error)
	Revoke(ctx context.Context, token string) error
	RevokeByKey(ctx context.Context, deviceID, userID string) error
}
