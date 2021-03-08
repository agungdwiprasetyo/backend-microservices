// Code generated by candi v1.4.0.

package graphqlhandler

import (
	"monorepo/services/master-service/internal/modules/acl/usecase"

	"pkg.agungdp.dev/candi/codebase/factory/types"
	"pkg.agungdp.dev/candi/codebase/interfaces"
)

// GraphQLHandler model
type GraphQLHandler struct {
	mw        interfaces.Middleware
	uc        usecase.ACLUsecase
	validator interfaces.Validator
}

// NewGraphQLHandler delivery
func NewGraphQLHandler(mw interfaces.Middleware, uc usecase.ACLUsecase, validator interfaces.Validator) *GraphQLHandler {
	return &GraphQLHandler{
		mw: mw, uc: uc, validator: validator,
	}
}

// RegisterMiddleware register resolver based on schema in "api/graphql/*" path
func (h *GraphQLHandler) RegisterMiddleware(mwGroup *types.MiddlewareGroup) {
	mwGroup.Add("AclQueryModule.hello", h.mw.GraphQLBearerAuth)
	mwGroup.Add("AclMutationModule.hello", h.mw.GraphQLBasicAuth)
}

// Query method
func (h *GraphQLHandler) Query() interface{} {
	return &queryResolver{root: h}
}

// Mutation method
func (h *GraphQLHandler) Mutation() interface{} {
	return &mutationResolver{root: h}
}

// Subscription method
func (h *GraphQLHandler) Subscription() interface{} {
	return &subscriptionResolver{root: h}
}