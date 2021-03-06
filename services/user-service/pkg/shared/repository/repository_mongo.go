// Code generated by candi v1.3.1. DO NOT EDIT.

package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	authrepo "monorepo/services/user-service/internal/modules/auth/repository"
	memberrepo "monorepo/services/user-service/internal/modules/member/repository"
)

// RepoMongo uow
type RepoMongo struct {
	readDB, writeDB *mongo.Database

	// register all repository from modules
	AuthRepo   authrepo.AuthRepository
	MemberRepo memberrepo.MemberRepository
}

var globalRepoMongo = new(RepoMongo)

// setSharedRepoMongo set the global singleton "RepoMongo" implementation
func setSharedRepoMongo(readDB, writeDB *mongo.Database) {
	globalRepoMongo = &RepoMongo{
		readDB: readDB, writeDB: writeDB,
		AuthRepo:   authrepo.NewAuthRepoMongo(readDB, writeDB),
		MemberRepo: memberrepo.NewMemberRepoMongo(readDB, writeDB),
	}
}

// GetSharedRepoMongo returns the global singleton "RepoMongo" implementation
func GetSharedRepoMongo() *RepoMongo {
	return globalRepoMongo
}
