// Code generated by candi v1.8.17.

package graphqlhandler

import (
	"context"

	"monorepo/services/user-service/internal/modules/member/domain"
	shareddomain "monorepo/services/user-service/pkg/shared/domain"

	"pkg.agungdp.dev/candi/tracer"
)

type mutationResolver struct {
	root *GraphQLHandler
}

// SaveMember resolver
func (m *mutationResolver) SaveMember(ctx context.Context, input struct{ Data shareddomain.Member }) (ok string, err error) {
	trace := tracer.StartTrace(ctx, "MemberDeliveryGraphQL:SaveMember")
	defer trace.Finish()
	ctx = trace.Context()

	// tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using GraphQLBearerAuth in middleware for this resolver

	if err := m.root.uc.SaveMember(ctx, &input.Data); err != nil {
		return ok, err
	}
	return "Success", nil
}

// SaveMember resolver
func (m *mutationResolver) Register(ctx context.Context, input struct{ Data domain.RegisterPayload }) (ok string, err error) {
	trace := tracer.StartTrace(ctx, "MemberDeliveryGraphQL:Register")
	defer trace.Finish()
	ctx = trace.Context()

	// tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using GraphQLBearerAuth in middleware for this resolver

	if err := m.root.uc.Register(ctx, &input.Data); err != nil {
		return ok, err
	}
	return "Success", nil
}
