// Code generated by candi v1.3.3.

package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"pkg.agungdwiprasetyo.com/candi/tracer"
)

type orderRepoMongo struct {
	readDB, writeDB *mongo.Database
}

// NewOrderRepoMongo mongo repo constructor
func NewOrderRepoMongo(readDB, writeDB *mongo.Database) OrderRepository {
	return &orderRepoMongo{
		readDB, writeDB,
	}
}

func (r *orderRepoMongo) FindHello(ctx context.Context) (string, error) {
	trace := tracer.StartTrace(ctx, "OrderRepoMongo:FindHello")
	defer trace.Finish()

	return "Hello from repo mongo layer", nil
}