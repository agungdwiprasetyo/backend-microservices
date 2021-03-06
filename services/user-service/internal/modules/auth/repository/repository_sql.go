// Code generated by candi v1.3.3.

package repository

import (
	"context"
	"database/sql"

	"pkg.agungdp.dev/candi/tracer"
)

type authRepoSQL struct {
	readDB, writeDB *sql.DB
	tx              *sql.Tx
}

// NewAuthRepoSQL mongo repo constructor
func NewAuthRepoSQL(readDB, writeDB *sql.DB, tx *sql.Tx) AuthRepository {
	return &authRepoSQL{
		readDB, writeDB, tx,
	}
}

func (r *authRepoSQL) FindHello(ctx context.Context) (string, error) {
	trace := tracer.StartTrace(ctx, "AuthRepoSQL:FindHello")
	defer trace.Finish()

	return "Hello from repo sql layer", nil
}
