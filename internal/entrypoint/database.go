package entrypoint

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Database interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}
