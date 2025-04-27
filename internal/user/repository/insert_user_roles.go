package user_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/Flak34/crowd-api/internal/user/model"
	"github.com/huandu/go-sqlbuilder"
)

func (r *Repository) InsertUserRoles(ctx context.Context, db entrypoint.Database, userID int, roles []user_model.Role) error {
	if len(roles) == 0 {
		return nil
	}
	ib := sqlbuilder.PostgreSQL.NewInsertBuilder()
	ib.InsertInto("person_role").
		Cols("person_id", "role_id")
	for _, role := range roles {
		ib.Values(userID, role.ID)
	}
	query, args := ib.Build()
	rows, _ := db.Query(ctx, query, args...)
	rows.Close()
	err := rows.Err()
	if err != nil {
		return err
	}
	return nil
}
