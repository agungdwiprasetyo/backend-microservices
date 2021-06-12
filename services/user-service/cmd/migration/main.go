package main

import (
	"context"
	"fmt"
	"log"
	"monorepo/services/user-service/internal/modules/member/repository"
	"monorepo/services/user-service/pkg/shared/domain"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pkg.agungdp.dev/candi/candihelper"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/config/database"
	"pkg.agungdp.dev/candi/config/env"
)

var databaseConnection, dbname, userID string

func main() {

	ctx := context.Background()
	env.Load("caterpillar")

	db := database.InitMongoDB(ctx)
	createMongoIndex(ctx, db)
	createSeed(ctx, db)
	migratePostgres(database.InitSQLDatabase())
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

func migratePostgres(db interfaces.SQLDatabase) {
	gormWrite, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db.WriteDB(),
	}), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := gormWrite.AutoMigrate(
		&domain.Member{},
	); err != nil {
		log.Fatal(err)
	}
	log.Printf("\x1b[32;1mMigration to \"%s\" suceess\x1b[0m\n", candihelper.MaskingPasswordURL(env.BaseEnv().DbSQLWriteDSN))
}
