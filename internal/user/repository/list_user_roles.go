package user_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/Flak34/crowd-api/internal/user/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/samber/lo"
)

func (r *Repository) ListUserRoles(ctx context.Context, db entrypoint.Database, userID int) ([]user_model.Role, error) {
	query := `SELECT id, name 
			  FROM role JOIN 
    			(SELECT role_id FROM person_role WHERE person_id = $1) role_ids
    		  ON id = role_ids.role_id`
	var roles []RoleTable
	err := pgxscan.Select(ctx, db, &roles, query, userID)
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
