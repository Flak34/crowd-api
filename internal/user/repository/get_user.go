package user_repository

import (
	"context"
	"database/sql"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/Flak34/crowd-api/internal/errors/storage_errors"
	"github.com/Flak34/crowd-api/internal/user/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
)

func (r *Repository) GetUser(ctx context.Context, db entrypoint.Database, email string) (user_model.User, error) {
	query := `SELECT id, pass_hash, email FROM person WHERE email = $1`
	var user PersonTable
	err := pgxscan.Get(ctx, db, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user_model.User{}, &storage_errors.ErrEntityNotFound{Entity: storage_errors.EntityUser}
		}
		return user_model.User{}, err
	}
	return user_model.User{
		ID:       user.ID,
		Email:    user.Email,
		PassHash: user.PassHash,
	}, nil
}
