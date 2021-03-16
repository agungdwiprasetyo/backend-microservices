// Code generated by candi v1.4.0.

package graphqlhandler

import (
	"context"
	"monorepo/services/master-service/internal/modules/acl/domain"

	"pkg.agungdp.dev/candi/tracer"
)

type queryResolver struct {
	root *GraphQLHandler
}

// GetAllRole resolver
func (q *queryResolver) GetAllRole(ctx context.Context, input struct{ Filter *CommonFilter }) (resp RoleResult, err error) {
	trace := tracer.StartTrace(ctx, "AclDeliveryGraphQL:GetAllRole")
	defer trace.Finish()
	ctx = trace.Context()

	if input.Filter == nil {
		input.Filter = new(CommonFilter)
	}
	data, meta, err := q.root.uc.GetAllRole(ctx, domain.RoleListFilter{Filter: input.Filter.ToSharedFilter()})
	if err != nil {
		return resp, err
	}

	resp.Meta = meta
	resp.Data = data
	return
}
