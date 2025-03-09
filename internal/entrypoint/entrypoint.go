package entrypoint

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

type entrypoint struct {
	pool connectionPool
}

type connectionPool interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

func New(pool connectionPool) Entrypoint {
	e := &entrypoint{
		pool: pool,
	}
	return e
}

type Entrypoint interface {
	GetDB() Database
	TxWrapper(ctx context.Context, fn func(ctx context.Context, tx pgx.Tx) error) (err error)
}

func (e *entrypoint) GetDB() Database {
	return e.pool
}

func (e *entrypoint) TxWrapper(ctx context.Context, fn func(ctx context.Context, tx pgx.Tx) error) (err error) {
	tx, err := e.pool.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "begin transaction")
	}
	defer func() {
		exc := recover()
		if exc != nil {
			err = errors.Errorf("panic %v", exc)
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				err = errors.Wrapf(rollbackErr, "rollback transaction after panic: %s", err.Error())
			}
		}
	}()
	err = fn(ctx, tx)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			err = errors.Wrapf(rollbackErr, "rollback transaction after error: %s", err.Error())
		}
		return err
	}
	return tx.Commit(ctx)
}
