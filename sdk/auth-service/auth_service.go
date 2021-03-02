package authservice

import (
	"context"

	"pkg.agungdp.dev/candi/candishared"
)

// AuthService abstract interface
type AuthService interface {
	ValidateToken(ctx context.Context, token string) (*candishared.TokenClaim, error)
	GenerateToken(ctx context.Context, req PayloadGenerateToken) (token string, err error)
}
