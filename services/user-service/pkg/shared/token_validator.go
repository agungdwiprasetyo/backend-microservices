// Code generated by candi v1.3.1.

package shared

import (
	"context"

	"pkg.agungdwiprasetyo.com/candi/candishared"
)

// DefaultTokenValidator for token validator
type DefaultTokenValidator struct {
}

// ValidateToken implement TokenValidator
func (v *DefaultTokenValidator) ValidateToken(ctx context.Context, token string) (*candishared.TokenClaim, error) {
	return &candishared.TokenClaim{}, nil
}
