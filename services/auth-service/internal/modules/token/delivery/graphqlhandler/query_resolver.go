// Code generated by candi v1.8.17.

package graphqlhandler

import (
	"context"

	shareddomain "monorepo/services/auth-service/pkg/shared/domain"

	"pkg.agungdp.dev/candi/tracer"
)

type queryResolver struct {
	root *GraphQLHandler
}

// GetDetailToken resolver
func (q *queryResolver) GetDetailToken(ctx context.Context, input struct{ ID string }) (data shareddomain.Token, err error) {
	trace := tracer.StartTrace(ctx, "TokenDeliveryGraphQL:GetDetailToken")
	defer trace.Finish()
	ctx = trace.Context()

	// tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using GraphQLBearerAuth in middleware for this resolver

	return
}
