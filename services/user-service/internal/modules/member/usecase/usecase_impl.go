// Code generated by candi v1.3.1.

package usecase

import (
	"context"

	shareddomain "monorepo/services/user-service/pkg/shared/domain"
	"monorepo/services/user-service/pkg/shared/repository"

	"pkg.agungdp.dev/candi/codebase/factory/dependency"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
)

type memberUsecaseImpl struct {
	cache interfaces.Cache

	repoMongo *repository.RepoMongo
}

// NewMemberUsecase usecase impl constructor
func NewMemberUsecase(deps dependency.Dependency) MemberUsecase {
	return &memberUsecaseImpl{
		cache: deps.GetRedisPool().Cache(),

		repoMongo: repository.GetSharedRepoMongo(),
	}
}

func (uc *memberUsecaseImpl) Hello(ctx context.Context) (msg string) {
	trace := tracer.StartTrace(ctx, "MemberUsecase:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	return
}

func (uc *memberUsecaseImpl) Save(ctx context.Context, data *shareddomain.Member) (err error) {
	trace := tracer.StartTrace(ctx, "AppsUsecase:FindAll")
	defer trace.Finish()
	ctx = trace.Context()

	return uc.repoMongo.MemberRepo.Save(ctx, data)
}

func (uc *memberUsecaseImpl) GetMemberByID(ctx context.Context, id string) (data shareddomain.Member, err error) {
	trace := tracer.StartTrace(ctx, "AppsUsecase:GetMemberByID")
	defer trace.Finish()
	ctx = trace.Context()

	data.ID = id
	err = uc.repoMongo.MemberRepo.Find(ctx, &data)
	return
}
