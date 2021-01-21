// Code generated by candi v1.3.3. DO NOT EDIT.

package repository

import (
	"sync"

	"pkg.agungdwiprasetyo.com/candi/codebase/factory/dependency"
)

var (
	once sync.Once
)

// SetSharedRepository set the global singleton "RepoSQL" and "RepoMongo" implementation
func SetSharedRepository(deps dependency.Dependency) {
	once.Do(func() {
		// setSharedRepoSQL(deps.GetSQLDatabase().ReadDB(), deps.GetSQLDatabase().WriteDB())
		setSharedRepoMongo(deps.GetMongoDatabase().ReadDB(), deps.GetMongoDatabase().WriteDB())
	})
}
