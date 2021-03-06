// Code generated by candi v1.8.18.

package repository

import (
	"context"
	"monorepo/services/master-service/internal/modules/acl/domain"
	shareddomain "monorepo/services/master-service/pkg/shared/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"pkg.agungdp.dev/candi/candihelper"
	"pkg.agungdp.dev/candi/tracer"
)

type aclRepoMongo struct {
	readDB, writeDB *mongo.Database
	collection      string
}

// NewACLRepoMongo mongo repo constructor
func NewACLRepoMongo(readDB, writeDB *mongo.Database) ACLRepository {
	return &aclRepoMongo{
		readDB, writeDB, shareddomain.ACL{}.CollectionName(),
	}
}

func (r *aclRepoMongo) FetchAll(ctx context.Context, filter domain.ACLFilter) (data []shareddomain.ACL, err error) {
	trace := tracer.StartTrace(ctx, "ACLRepoMongo:FetchAll")
	defer trace.Finish()
	defer func() { trace.SetError(err) }()
	ctx = trace.Context()

	where := bson.M{}
	if filter.AppsID != "" {
		where["appsId"] = filter.AppsID
	}
	if filter.UserID != "" {
		where["userId"] = filter.UserID
	}
	trace.SetTag("query", where)

	findOptions := options.Find()
	if len(filter.OrderBy) > 0 {
		findOptions.SetSort(filter)
	}

	if !filter.ShowAll {
		findOptions.SetLimit(int64(filter.Limit))
		findOptions.SetSkip(int64(filter.Offset))
	}
	cur, err := r.readDB.Collection(r.collection).Find(ctx, where, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	cur.All(ctx, &data)
	trace.Log("result", data)
	return
}

func (r *aclRepoMongo) FindByUserID(ctx context.Context, userID string) (data []shareddomain.ACL, err error) {
	trace := tracer.StartTrace(ctx, "AclRepoMongo:FindByUserID")
	ctx = trace.Context()
	defer func() { tracer.Log(ctx, "results", data); trace.Finish() }()

	cur, err := r.readDB.Collection(r.collection).Find(ctx, bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var acl shareddomain.ACL
		err := cur.Decode(&acl)
		if err != nil {
			return data, err
		}
		data = append(data, acl)
	}

	return
}

func (r *aclRepoMongo) Find(ctx context.Context, data *shareddomain.ACL) (err error) {
	trace := tracer.StartTrace(ctx, "AclRepoMongo:Find")
	defer trace.Finish()

	bsonWhere := make(bson.M)
	if data.ID != "" {
		bsonWhere["_id"] = data.ID
	}
	if data.UserID != "" {
		bsonWhere["userId"] = data.UserID
	}
	if data.RoleID != "" {
		bsonWhere["roleId"] = data.RoleID
	}
	trace.SetTag("query", bsonWhere)

	return r.readDB.Collection(r.collection).FindOne(ctx, bsonWhere).Decode(data)
}

func (r *aclRepoMongo) Save(ctx context.Context, data *shareddomain.ACL) (err error) {
	trace := tracer.StartTrace(ctx, "AclRepoMongo:Save")
	defer func() { trace.SetError(err); trace.Finish() }()
	ctx = trace.Context()
	tracer.Log(ctx, "data", data)

	data.ModifiedAt = time.Now()
	if data.ID == "" {
		data.ID = primitive.NewObjectID().Hex()
		data.CreatedAt = time.Now()
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

func (r *aclRepoMongo) Delete(ctx context.Context, data *shareddomain.ACL) (err error) {
	trace := tracer.StartTrace(ctx, "AclRepoMongo:Delete")
	defer func() { trace.SetError(err); trace.Finish() }()
	ctx = trace.Context()
	tracer.Log(ctx, "data", data)

	_, err = r.writeDB.Collection(r.collection).DeleteOne(ctx, bson.M{"_id": data.ID})
	return
}
