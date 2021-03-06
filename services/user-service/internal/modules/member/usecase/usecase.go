// Code generated by candi v1.8.17.

package usecase

import (
	"context"

	"monorepo/services/user-service/internal/modules/member/domain"
	shareddomain "monorepo/services/user-service/pkg/shared/domain"

	"pkg.agungdp.dev/candi/candishared"
)

// MemberUsecase abstraction
type MemberUsecase interface {
	GetAllMember(ctx context.Context, filter candishared.Filter) (data []shareddomain.Member, meta candishared.Meta, err error)
	GetDetailMember(ctx context.Context, id string) (data shareddomain.Member, err error)
	SaveMember(ctx context.Context, data *shareddomain.Member) (err error)
	Register(ctx context.Context, data *domain.RegisterPayload) (err error)
	AutoGenerateMember(ctx context.Context) error
}
