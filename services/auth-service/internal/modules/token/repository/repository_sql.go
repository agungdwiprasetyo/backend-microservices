// Code generated by candi v1.3.3.

package repository

import (
	"context"
	"database/sql"

	"pkg.agungdp.dev/candi/tracer"
)

type tokenRepoSQL struct {
	readDB, writeDB *sql.DB
	tx              *sql.Tx
}

// NewTokenRepoSQL mongo repo constructor
func NewTokenRepoSQL(readDB, writeDB *sql.DB, tx *sql.Tx) TokenRepository {
	return &tokenRepoSQL{
		readDB, writeDB, tx,
	}
}

func (r *tokenRepoSQL) FindHello(ctx context.Context) (string, error) {
	trace := tracer.StartTrace(ctx, "TokenRepoSQL:FindHello")
	defer trace.Finish()

	return "Hello from repo sql layer", nil
}
