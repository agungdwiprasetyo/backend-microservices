// Code generated by candi v1.3.3.

package master

import (
	"monorepo/services/order-service/internal/modules/master/delivery/graphqlhandler"
	"monorepo/services/order-service/internal/modules/master/delivery/grpchandler"

	// "monorepo/services/order-service/internal/modules/master/delivery/resthandler"
	// "monorepo/services/order-service/internal/modules/master/delivery/workerhandler"
	"monorepo/services/order-service/pkg/shared/usecase"

	"pkg.agungdwiprasetyo.com/candi/codebase/factory/dependency"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/types"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
)

const (
	moduleName types.Module = "Master"
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
	// mod.restHandler = resthandler.NewRestHandler(deps.GetMiddleware(), usecaseUOW.Master(), deps.GetValidator())
	mod.grpcHandler = grpchandler.NewGRPCHandler(deps.GetMiddleware(), usecaseUOW.Master(), deps.GetValidator())
	mod.graphqlHandler = graphqlhandler.NewGraphQLHandler(deps.GetMiddleware(), usecaseUOW.Master(), deps.GetValidator())

	mod.workerHandlers = map[types.Worker]interfaces.WorkerHandler{
		// types.Kafka:           workerhandler.NewKafkaHandler(usecaseUOW.Master(), deps.GetValidator()),
		// types.Scheduler:       workerhandler.NewCronHandler(usecaseUOW.Master(), deps.GetValidator()),
		// types.RedisSubscriber: workerhandler.NewRedisHandler(usecaseUOW.Master(), deps.GetValidator()),
		// types.TaskQueue:       workerhandler.NewTaskQueueHandler(usecaseUOW.Master(), deps.GetValidator()),
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
