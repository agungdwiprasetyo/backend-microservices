package main

import (
	"context"
	aclrepo "monorepo/services/master-service/internal/modules/acl/repository"
	appsrepo "monorepo/services/master-service/internal/modules/apps/repository"
	"monorepo/services/master-service/pkg/shared/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/logger"
)

func migrateMongo(ctx context.Context, db interfaces.MongoDatabase) {
	createMongoIndex(ctx, db)
	createSeed(ctx, db)
}

func createMongoIndex(ctx context.Context, db interfaces.MongoDatabase) {
	indexACL := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "userId", Value: 1},
				{Key: "appsId", Value: 1},
				{Key: "roleId", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
	indexViewACL := db.WriteDB().Collection(domain.ACL{}.CollectionName()).Indexes()
	if _, err := indexViewACL.CreateMany(ctx, indexACL); err != nil {
		panic(err)
	}

	indexRole := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "appsId", Value: 1},
				{Key: "code", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
	indexViewRole := db.WriteDB().Collection(domain.Role{}.CollectionName()).Indexes()
	if _, err := indexViewRole.CreateMany(ctx, indexRole); err != nil {
		panic(err)
	}

	indexApps := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "code", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
	indexViewApps := db.WriteDB().Collection(domain.Apps{}.CollectionName()).Indexes()
	if _, err := indexViewApps.CreateMany(ctx, indexApps); err != nil {
		panic(err)
	}
	logger.LogGreen("Success create index")
}

func createSeed(ctx context.Context, db interfaces.MongoDatabase) {
	aclRepo := aclrepo.NewACLRepoMongo(db.ReadDB(), db.WriteDB())
	roleRepo := aclrepo.NewRoleRepoMongo(db.ReadDB(), db.WriteDB())
	appsRepo := appsrepo.NewAppsRepoMongo(db.ReadDB(), db.WriteDB())
	permissionRepo := appsrepo.NewPermissionRepoMongo(db.ReadDB(), db.WriteDB())

	appsData := domain.Apps{Code: "caterpillar", Name: "Caterpillar"}
	if err := appsRepo.Save(ctx, &appsData); err != nil {
		panic(err)
	}

	permissionData := []domain.Permission{
		{AppsID: appsData.ID, Code: "apps", Name: "Application", Childs: []domain.Permission{
			{AppsID: appsData.ID, Code: "getAllApps", Name: "Get All Apps", Childs: []domain.Permission{
				{AppsID: appsData.ID, Code: "addApps", Name: "Add Apps"},
				{AppsID: appsData.ID, Code: "getDetailApp", Name: "Get Detail App"},
			}},
		}},
		{AppsID: appsData.ID, Code: "acl", Name: "Access Control List", Childs: []domain.Permission{
			{AppsID: appsData.ID, Code: "getAllRole", Name: "Get All Role", Childs: []domain.Permission{
				{AppsID: appsData.ID, Code: "addRole", Name: "Add Role"},
				{AppsID: appsData.ID, Code: "getDetailRole", Name: "Get Detail Role"},
			}},
			{AppsID: appsData.ID, Code: "grantUser", Name: "Grant User Role"},
			{AppsID: appsData.ID, Code: "revokeUser", Name: "Revoke User Role"},
		}},
	}

	roleData := domain.Role{AppsID: appsData.ID, Code: "superadmin", Name: "Super Admin", Permissions: make(map[string]string)}
	for _, permission := range permissionData {
		for _, tree := range fetchAllPermission(permission) {
			roleData.Permissions[tree.Code] = tree.ID
			if err := permissionRepo.Save(ctx, &tree); err != nil {
				panic(err)
			}
		}
	}

	if err := roleRepo.Save(ctx, &roleData); err != nil {
		panic(err)
	}
	if err := aclRepo.Save(ctx, &domain.ACL{
		UserID: "random_00001",
		AppsID: appsData.ID,
		RoleID: roleData.ID,
	}); err != nil {
		panic(err)
	}
	logger.LogGreen("Success create data seed")
}

func fetchAllPermission(data domain.Permission) (result []domain.Permission) {
	data.ID = primitive.NewObjectID().Hex()
	result = append(result, data)
	for _, perm := range data.Childs {
		perm.ParentID = data.ID
		result = append(result, fetchAllPermission(perm)...)
	}
	return
}
