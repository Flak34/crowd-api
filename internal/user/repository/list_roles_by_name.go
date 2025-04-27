package user_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/Flak34/crowd-api/internal/user/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/samber/lo"
)

func (r *Repository) ListRolesByName(ctx context.Context, db entrypoint.Database, roleNames ...string) ([]user_model.Role, error) {
	query := `SELECT id, name FROM role WHERE name = ANY ($1)`
	var roles []RoleTable
	err := pgxscan.Select(ctx, db, &roles, query, roleNames)
	if err != nil {
		return nil, err
	}
	return lo.Map(roles, func(role RoleTable, _ int) user_model.Role {
		return user_model.Role{
			Name: role.Name,
			ID:   role.ID,
		}
	}), nil
}
