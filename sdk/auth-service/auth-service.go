package authservice

import (
	"context"

	"pkg.agungdwiprasetyo.com/candi/candishared"
)

// AuthService abstract interface
type AuthService interface {
	ValidateToken(ctx context.Context, token string) (*candishared.TokenClaim, error)
}
