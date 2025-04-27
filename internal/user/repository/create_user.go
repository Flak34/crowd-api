package user_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/Flak34/crowd-api/internal/errors/storage_errors"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

const (
	personEmailUqConstraintName = "person_email_key"
)

func (r *Repository) CreateUser(ctx context.Context, db entrypoint.Database, email string, passHash []byte) (int, error) {
	var query = `INSERT INTO person(email, pass_hash) VALUES ($1, $2) RETURNING id`
	var id int
	err := pgxscan.Get(ctx, db, &id, query, email, passHash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == personEmailUqConstraintName {
			return id, &storage_errors.ErrDuplicateKey{UniqueConstraint: storage_errors.PersonEmailUqConstraint}
		}
		return id, err
	}
	return id, nil
}
