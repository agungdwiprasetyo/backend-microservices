// Code generated by candi v1.3.1.

package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"pkg.agungdwiprasetyo.com/candi/tracer"
)

type storageRepoMongo struct {
	readDB, writeDB *mongo.Database
}

// NewStorageRepoMongo mongo repo constructor
func NewStorageRepoMongo(readDB, writeDB *mongo.Database) StorageRepository {
	return &storageRepoMongo{
		readDB, writeDB,
	}
}

func (r *storageRepoMongo) FindHello(ctx context.Context) (string, error) {
	trace := tracer.StartTrace(ctx, "StorageRepoMongo:FindHello")
	defer trace.Finish()

	return "Hello from repo mongo layer", nil
}