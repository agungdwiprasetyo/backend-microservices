package event

import (
	"agungdwiprasetyo.com/backend-microservices/internal/factory/base"
	"agungdwiprasetyo.com/backend-microservices/internal/factory/constant"
	"agungdwiprasetyo.com/backend-microservices/internal/factory/interfaces"
	"agungdwiprasetyo.com/backend-microservices/internal/services/line-chatbot/modules/event/delivery/graphqlhandler"
	"agungdwiprasetyo.com/backend-microservices/internal/services/line-chatbot/modules/event/delivery/grpchandler"
	"agungdwiprasetyo.com/backend-microservices/internal/services/line-chatbot/modules/event/repository"
	"agungdwiprasetyo.com/backend-microservices/internal/services/line-chatbot/modules/event/usecase"
	"agungdwiprasetyo.com/backend-microservices/pkg/helper"
)

const (
	// Event service name
	Event constant.Module = "Event"
)

// Module model
type Module struct {
	graphqlHandler *graphqlhandler.GraphQLHandler
	grpcHandler    *grpchandler.GRPCHandler
}

// NewModule module constructor
func NewModule(params *base.ModuleParam) *Module {
	repo := repository.NewRepoMongo(params.Config.MongoWrite)
	uc := usecase.NewEventUsecase(repo)

	var mod Module
	mod.graphqlHandler = graphqlhandler.NewGraphQLHandler(params.Middleware, uc)
	mod.grpcHandler = grpchandler.NewGRPCHandler(uc)
	return &mod
}

// RestHandler method
func (m *Module) RestHandler(version string) (d interfaces.EchoRestHandler) {
	switch version {
	case helper.V1:
		d = nil
	case helper.V2:
		d = nil // TODO versioning
	}
	return
}

// GRPCHandler method
func (m *Module) GRPCHandler() interfaces.GRPCHandler {
	return m.grpcHandler
}

// GraphQLHandler method
func (m *Module) GraphQLHandler() (name string, resolver interface{}) {
	return string(Event), m.graphqlHandler
}

// SubscriberHandler method
func (m *Module) SubscriberHandler(subsType constant.Subscriber) interfaces.SubscriberHandler {
	return nil
}

// Name get module name
func (m *Module) Name() constant.Module {
	return Event
}