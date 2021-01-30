// Code generated by candi v1.3.1.

package member

import (
	"monorepo/services/user-service/internal/modules/member/delivery/graphqlhandler"
	// "monorepo/services/user-service/internal/modules/member/delivery/grpchandler"
	// "monorepo/services/user-service/internal/modules/member/delivery/resthandler"
	"monorepo/services/user-service/internal/modules/member/delivery/workerhandler"
	"monorepo/services/user-service/pkg/shared/usecase"

	"pkg.agungdwiprasetyo.com/candi/codebase/factory/dependency"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/types"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
)

const (
	moduleName types.Module = "Member"
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
	// mod.restHandler = resthandler.NewRestHandler(deps.GetMiddleware(), usecaseUOW.Member(), deps.GetValidator())
	// mod.grpcHandler = grpchandler.NewGRPCHandler(deps.GetMiddleware(), usecaseUOW.Member(), deps.GetValidator())
	mod.graphqlHandler = graphqlhandler.NewGraphQLHandler(deps.GetMiddleware(), usecaseUOW.Member(), deps.GetValidator())

	mod.workerHandlers = map[types.Worker]interfaces.WorkerHandler{
		types.Kafka: workerhandler.NewKafkaHandler(usecaseUOW.Member(), deps.GetValidator()),
		// types.Scheduler:       workerhandler.NewCronHandler(usecaseUOW.Member(), deps.GetValidator()),
		// types.RedisSubscriber: workerhandler.NewRedisHandler(usecaseUOW.Member(), deps.GetValidator()),
		// types.TaskQueue:       workerhandler.NewTaskQueueHandler(usecaseUOW.Member(), deps.GetValidator()),
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
