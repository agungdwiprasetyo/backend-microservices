// Code generated by candi v1.8.17.

package repository

import (
	"context"

	shareddomain "monorepo/services/user-service/pkg/shared/domain"

	"pkg.agungdp.dev/candi/candishared"
)

// MemberRepository abstract interface
type MemberRepository interface {
	FetchAll(ctx context.Context, filter *candishared.Filter) ([]shareddomain.Member, error)
	Count(ctx context.Context, filter *candishared.Filter) int
	Find(ctx context.Context, data *shareddomain.Member) error
	Save(ctx context.Context, data *shareddomain.Member) error
}
