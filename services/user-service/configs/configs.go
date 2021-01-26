// Code generated by candi v1.3.1.

package configs

import (
	"context"

	authservice "monorepo/sdk/auth-service"
	"monorepo/services/user-service/pkg/shared"
	"monorepo/services/user-service/pkg/shared/repository"
	"monorepo/services/user-service/pkg/shared/usecase"

	"pkg.agungdwiprasetyo.com/candi/candihelper"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/dependency"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/types"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
	"pkg.agungdwiprasetyo.com/candi/config"
	"pkg.agungdwiprasetyo.com/candi/config/broker"
	"pkg.agungdwiprasetyo.com/candi/config/database"
	"pkg.agungdwiprasetyo.com/candi/middleware"
	"pkg.agungdwiprasetyo.com/candi/validator"
)

// LoadConfigs load selected dependency configuration in this service
func LoadConfigs(baseCfg *config.Config) (deps dependency.Dependency) {

	var sharedEnv shared.Environment
	candihelper.MustParseEnv(&sharedEnv)
	shared.SetEnv(sharedEnv)

	baseCfg.LoadFunc(func(ctx context.Context) []interfaces.Closer {
		brokerDeps := broker.InitBrokers(
			types.Kafka,
		)
		redisDeps := database.InitRedis()
		// sqlDeps := database.InitSQLDatabase()
		mongoDeps := database.InitMongoDB(ctx)

		authService := authservice.NewAuthServiceGRPC(sharedEnv.AuthServiceHost, sharedEnv.AuthServiceKey)

		// inject all service dependencies
		// See all option in dependency package
		deps = dependency.InitDependency(
			dependency.SetMiddleware(middleware.NewMiddleware(authService)),
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
