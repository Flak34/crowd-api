package project_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	model "github.com/Flak34/crowd-api/internal/project/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
)

func (r *Repository) ListProjects(ctx context.Context, db entrypoint.Database, projectIDs ...int) ([]model.Project, error) {
	var query = `
		SELECT 
		    id, 
		    creator_id,
		    description,
		    name,
		    instruction,
		    task_config, 
		    target_overlap, 
		    tasks_per_user, 
		    created_at,
		    annotator_time_limit
		FROM project
		WHERE id = ANY($1::INTEGER[])`
	var projects []*ProjectTable
	err := pgxscan.Select(ctx, db, &projects, query, pgtype.FlatArray[int](projectIDs))
	if err != nil {
		return []model.Project{}, err
	}
	return lo.FilterMap(projects, func(p *ProjectTable, _ int) (model.Project, bool) {
		if p == nil {
			return model.Project{}, false
		}
		return mapProjectTableToModel(*p), true
	}), nil
}
