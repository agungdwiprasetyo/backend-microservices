// Code generated by candi v1.3.3.

package graphqlhandler

import (
	"monorepo/services/order-service/internal/modules/order/usecase"

	"pkg.agungdp.dev/candi/codebase/factory/types"
	"pkg.agungdp.dev/candi/codebase/interfaces"
)

// GraphQLHandler model
type GraphQLHandler struct {
	mw        interfaces.Middleware
	uc        usecase.OrderUsecase
	validator interfaces.Validator
}

// NewGraphQLHandler delivery
func NewGraphQLHandler(mw interfaces.Middleware, uc usecase.OrderUsecase, validator interfaces.Validator) *GraphQLHandler {
	return &GraphQLHandler{
		mw: mw, uc: uc, validator: validator,
	}
}

// RegisterMiddleware register resolver based on schema in "api/graphql/*" path
func (h *GraphQLHandler) RegisterMiddleware(mwGroup *types.MiddlewareGroup) {
	mwGroup.Add("OrderQueryModule.hello", h.mw.GraphQLBearerAuth)
	mwGroup.Add("OrderMutationModule.hello", h.mw.GraphQLBasicAuth)
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
