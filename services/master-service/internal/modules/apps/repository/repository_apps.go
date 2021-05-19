// Code generated by candi v1.8.18.

package repository

import (
	"context"
	"errors"
	"monorepo/services/master-service/internal/modules/apps/domain"
	shareddomain "monorepo/services/master-service/pkg/shared/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"pkg.agungdp.dev/candi/candihelper"
	"pkg.agungdp.dev/candi/tracer"
)

type appsRepoMongo struct {
	readDB, writeDB *mongo.Database
	collection      string
}

// NewAppsRepoMongo mongo repo constructor
func NewAppsRepoMongo(readDB, writeDB *mongo.Database) AppsRepository {
	return &appsRepoMongo{
		readDB, writeDB, shareddomain.Apps{}.CollectionName(),
	}
}

func (r *appsRepoMongo) FetchAll(ctx context.Context, filter domain.FilterApps) (data []shareddomain.Apps, err error) {
	trace := tracer.StartTrace(ctx, "AppsRepoMongo:FetchAll")
	defer trace.Finish()
	defer func() { trace.SetError(err) }()
	ctx = trace.Context()

	where := bson.M{}
	if len(filter.IDs) > 0 {
		where["_id"] = bson.M{
			"$in": filter.IDs,
		}
	}

	findOptions := options.Find()
	if len(filter.OrderBy) > 0 {
		findOptions.SetSort(filter.OrderBy)
	}

	findOptions.SetLimit(int64(filter.Limit))
	findOptions.SetSkip(int64(filter.Offset))
	cur, err := r.readDB.Collection(r.collection).Find(ctx, where, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var app shareddomain.Apps
		err := cur.Decode(&app)
		if err != nil {
			return data, err
		}
		app.CreatedAt = app.CreatedAtDB.Format(time.RFC3339)
		app.ModifiedAt = app.ModifiedAtDB.Format(time.RFC3339)
		data = append(data, app)
	}

	return
}

func (r *appsRepoMongo) Find(ctx context.Context, data *shareddomain.Apps) (err error) {
	trace := tracer.StartTrace(ctx, "AppsRepoMongo:Find")
	defer trace.Finish()
	defer func() { trace.SetError(err) }()
	ctx = trace.Context()

	bsonWhere := make(bson.M)
	if data.ID != "" {
		bsonWhere["_id"] = data.ID
	}
	if data.Code != "" {
		bsonWhere["code"] = data.Code
	}
	if data.Name != "" {
		bsonWhere["name"] = data.Name
	}
	if len(bsonWhere) == 0 {
		return errors.New("Role not found")
	}
	trace.SetTag("query", bsonWhere)

	return r.readDB.Collection(r.collection).FindOne(ctx, bsonWhere).Decode(data)
}

func (r *appsRepoMongo) Count(ctx context.Context, filter domain.FilterApps) int64 {
	trace := tracer.StartTrace(ctx, "AppsRepoMongo:Count")
	defer trace.Finish()

	count, err := r.readDB.Collection(r.collection).CountDocuments(trace.Context(), bson.M{})
	trace.SetError(err)
	return count
}

func (r *appsRepoMongo) Save(ctx context.Context, data *shareddomain.Apps) (err error) {
	trace := tracer.StartTrace(ctx, "AppsRepoMongo:Count")
	defer trace.Finish()
	defer func() { trace.SetError(err) }()
	ctx = trace.Context()
	tracer.Log(ctx, "data", data)

	data.ModifiedAtDB = time.Now()
	if data.ID == "" {
		data.ID = primitive.NewObjectID().Hex()
		data.CreatedAtDB = time.Now()
		_, err = r.writeDB.Collection(r.collection).InsertOne(ctx, data)
	} else {
		opt := options.UpdateOptions{
			Upsert: candihelper.ToBoolPtr(true),
		}
		_, err = r.writeDB.Collection(r.collection).UpdateOne(ctx,
			bson.M{
				"_id": data.ID,
			},
			bson.M{
				"$set": data,
			}, &opt)
	}

	return
}
