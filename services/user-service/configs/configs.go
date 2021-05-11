// Code generated by candi v1.3.1.

package configs

import (
	"context"

	"monorepo/sdk"
	authservice "monorepo/sdk/auth-service"
	masterservice "monorepo/sdk/master-service"
	"monorepo/services/user-service/pkg/shared"
	"monorepo/services/user-service/pkg/shared/repository"
	"monorepo/services/user-service/pkg/shared/usecase"

	"pkg.agungdp.dev/candi/candihelper"
	"pkg.agungdp.dev/candi/codebase/factory/dependency"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/config"
	"pkg.agungdp.dev/candi/config/broker"
	"pkg.agungdp.dev/candi/config/database"
	"pkg.agungdp.dev/candi/middleware"
	"pkg.agungdp.dev/candi/validator"
)

// LoadConfigs load selected dependency configuration in this service
func LoadConfigs(baseCfg *config.Config) (deps dependency.Dependency) {

	var sharedEnv shared.Environment
	candihelper.MustParseEnv(&sharedEnv)
	shared.SetEnv(sharedEnv)

	baseCfg.LoadFunc(func(ctx context.Context) []interfaces.Closer {
		brokerDeps := broker.InitBrokers(
		// broker.SetKafka(broker.NewKafkaBroker()),
		// broker.SetRabbitMQ(broker.NewRabbitMQBroker()),
		)
		redisDeps := database.InitRedis()
		// sqlDeps := database.InitSQLDatabase()
		mongoDeps := database.InitMongoDB(ctx)

		authService := authservice.NewAuthServiceGRPC(sharedEnv.AuthServiceHost, sharedEnv.AuthServiceKey)
		masterService := masterservice.NewMasterServiceGRPC(sharedEnv.MasterServiceHost, sharedEnv.MasterServiceKey)
		sdk.SetGlobalSDK(
			sdk.SetAuthService(authService),
			sdk.SetMasterService(masterService),
		)

		// inject all service dependencies
		// See all option in dependency package
		deps = dependency.InitDependency(
			dependency.SetMiddleware(middleware.NewMiddleware(authService, masterService)),
			dependency.SetValidator(validator.NewValidator()),
			dependency.SetBroker(brokerDeps),
			dependency.SetRedisPool(redisDeps),
			// dependency.SetSQLDatabase(sqlDeps),
			dependency.SetMongoDatabase(mongoDeps),
			// ... add more dependencies
		)
		return []interfaces.Closer{ // throw back to base config for close connection when application shutdown
			brokerDeps,
			redisDeps,
			// sqlDeps,
			mongoDeps,
		}
	})

	repository.SetSharedRepository(deps)
	usecase.SetSharedUsecase(deps)

	return deps
}
