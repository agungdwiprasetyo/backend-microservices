// Code generated by candi v1.3.1.

package usecase

import (
	"context"
	shareddomain "monorepo/services/user-service/pkg/shared/domain"

	"pkg.agungdp.dev/candi/candishared"
)

// MemberUsecase abstraction
type MemberUsecase interface {
	// add method
	Hello(ctx context.Context) string
	Save(ctx context.Context, data *shareddomain.Member) (err error)
	GetAllMember(ctx context.Context, filter candishared.Filter) (data []shareddomain.Member, meta candishared.Meta, err error)
	GetMemberByID(ctx context.Context, id string) (data shareddomain.Member, err error)
}
