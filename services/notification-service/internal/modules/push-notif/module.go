// Code generated by candi v1.3.1.

package pushnotif

import (
	"monorepo/services/notification-service/internal/modules/push-notif/delivery/graphqlhandler"
	// "monorepo/services/notification-service/internal/modules/push-notif/delivery/grpchandler"
	// "monorepo/services/notification-service/internal/modules/push-notif/delivery/resthandler"
	"monorepo/services/notification-service/internal/modules/push-notif/delivery/workerhandler"
	"monorepo/services/notification-service/pkg/shared/usecase"

	"pkg.agungdwiprasetyo.com/candi/codebase/factory/dependency"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/types"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
)

const (
	moduleName types.Module = "PushNotif"
)

// Module model
type Module struct {
	restHandler    interfaces.EchoRestHandler
	grpcHandler    interfaces.GRPCHandler
	graphqlHandler interfaces.GraphQLHandler

	workerHandlers map[types.Worker]interfaces.WorkerHandler
}

// NewModule module constructor
func NewModule(deps dependency.Dependency) *Module {
	usecaseUOW := usecase.GetSharedUsecase()

	var mod Module
	// mod.restHandler = resthandler.NewRestHandler(deps.GetMiddleware(), usecaseUOW.PushNotif(), deps.GetValidator())
	// mod.grpcHandler = grpchandler.NewGRPCHandler(deps.GetMiddleware(), usecaseUOW.PushNotif(), deps.GetValidator())
	mod.graphqlHandler = graphqlhandler.NewGraphQLHandler(deps.GetMiddleware(), usecaseUOW.PushNotif(), deps.GetValidator())

	mod.workerHandlers = map[types.Worker]interfaces.WorkerHandler{
		// types.Kafka:           workerhandler.NewKafkaHandler(usecaseUOW.PushNotif(), deps.GetValidator()),
		// types.Scheduler:       workerhandler.NewCronHandler(usecaseUOW.PushNotif(), deps.GetValidator()),
		types.RedisSubscriber: workerhandler.NewRedisHandler(usecaseUOW.PushNotif(), deps.GetValidator()),
		// types.TaskQueue:       workerhandler.NewTaskQueueHandler(usecaseUOW.PushNotif(), deps.GetValidator()),
	}

	return &mod
}

// RestHandler method
func (m *Module) RestHandler() interfaces.EchoRestHandler {
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
