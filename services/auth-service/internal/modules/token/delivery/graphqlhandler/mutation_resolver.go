// Code generated by candi v1.8.17.

package graphqlhandler

import (
	"context"
	"monorepo/services/auth-service/internal/modules/token/domain"

	"pkg.agungdp.dev/candi/tracer"
)

type mutationResolver struct {
	root *GraphQLHandler
}

// RefreshToken resolver
func (m *mutationResolver) RefreshToken(ctx context.Context, input struct{ Token, RefreshToken string }) (res domain.ResponseGenerateToken, err error) {
	trace := tracer.StartTrace(ctx, "TokenDeliveryGraphQL:RefreshToken")
	defer trace.Finish()
	ctx = trace.Context()

	// tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using GraphQLBearerAuth in middleware for this resolver

	return m.root.uc.Refresh(ctx, input.Token, input.RefreshToken)
}
