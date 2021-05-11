package main

import (
	"context"
	"flag"
	"fmt"
	"monorepo/services/user-service/internal/modules/member/repository"
	"monorepo/services/user-service/pkg/shared/domain"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/config/database"
	"pkg.agungdp.dev/candi/config/env"
)

var databaseConnection, dbname, userID string

func main() {
	flag.StringVar(&databaseConnection, "dbconn", "", "Database connection target")
	flag.StringVar(&dbname, "dbname", "admin", "Database name for create index")
	flag.StringVar(&userID, "userid", "admin", "User ID from user-service")
	flag.Parse()

	if databaseConnection == "" {
		flag.Usage()
		os.Exit(1)
	}

	ctx := context.Background()
	env.SetEnv(env.Env{DbMongoWriteHost: databaseConnection, DbMongoReadHost: databaseConnection, DbMongoDatabaseName: dbname})

	db := database.InitMongoDB(ctx)
	createMongoIndex(ctx, db)
	createSeed(ctx, db)
}

func createMongoIndex(ctx context.Context, db interfaces.MongoDatabase) {
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "username", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
	indexView := db.WriteDB().Collection("members").Indexes()
	for _, idx := range indexes {
		if _, err := indexView.CreateOne(ctx, idx); err != nil {
			panic(err)
		}
	}
}

func createSeed(ctx context.Context, db interfaces.MongoDatabase) {

	memberRepo := repository.NewMemberRepoMongo(db.ReadDB(), db.WriteDB())
	member := domain.Member{Username: "admin", Password: "plain"}
	memberRepo.Save(ctx, &member)
	fmt.Println("userID:", member.ID)
}
