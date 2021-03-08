// Code generated by candi v1.4.0.

package acl

import (
	// "monorepo/services/master-service/internal/modules/acl/delivery/graphqlhandler"
	"monorepo/services/master-service/internal/modules/acl/delivery/grpchandler"
	"monorepo/services/master-service/internal/modules/acl/delivery/resthandler"
	"monorepo/services/master-service/internal/modules/acl/delivery/workerhandler"
	"monorepo/services/master-service/pkg/shared/usecase"

	"pkg.agungdp.dev/candi/codebase/factory/dependency"
	"pkg.agungdp.dev/candi/codebase/factory/types"
	"pkg.agungdp.dev/candi/codebase/interfaces"
)

const (
	moduleName types.Module = "ACL"
)

// Module model
type Module struct {
	restHandler    interfaces.RESTHandler
	grpcHandler    interfaces.GRPCHandler
	graphqlHandler interfaces.GraphQLHandler

	workerHandlers map[types.Worker]interfaces.WorkerHandler
}

// NewModule module constructor
func NewModule(deps dependency.Dependency) *Module {
	usecaseUOW := usecase.GetSharedUsecase()

	var mod Module
	mod.restHandler = resthandler.NewRestHandler(deps.GetMiddleware(), usecaseUOW.ACL(), deps.GetValidator())
	mod.grpcHandler = grpchandler.NewGRPCHandler(deps.GetMiddleware(), usecaseUOW.ACL(), deps.GetValidator())
	// mod.graphqlHandler = graphqlhandler.NewGraphQLHandler(deps.GetMiddleware(), usecaseUOW.ACL(), deps.GetValidator())

	mod.workerHandlers = map[types.Worker]interfaces.WorkerHandler{
		// types.Kafka:           workerhandler.NewKafkaHandler(usecaseUOW.ACL(), deps.GetValidator()),
		// types.Scheduler:       workerhandler.NewCronHandler(usecaseUOW.ACL(), deps.GetValidator()),
		types.RedisSubscriber: workerhandler.NewRedisHandler(usecaseUOW.ACL(), deps.GetValidator()),
		// types.TaskQueue:       workerhandler.NewTaskQueueHandler(usecaseUOW.ACL(), deps.GetValidator()),
	}

	return &mod
}

// RESTHandler method
func (m *Module) RESTHandler() interfaces.RESTHandler {
	return m.restHandler
}

// GRPCHandler method
func (m *Module) GRPCHandler() interfaces.GRPCHandler {
	return m.grpcHandler
}

// GraphQLHandler method
func (m *Module) GraphQLHandler() interfaces.GraphQLHandler {
	return m.graphqlHandler
}

// WorkerHandler method
func (m *Module) WorkerHandler(workerType types.Worker) interfaces.WorkerHandler {
	return m.workerHandlers[workerType]
}

// Name get module name
func (m *Module) Name() types.Module {
	return moduleName
}