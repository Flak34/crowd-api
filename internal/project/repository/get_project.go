package project_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/Flak34/crowd-api/internal/errors/storage_errors"
	model "github.com/Flak34/crowd-api/internal/project/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (r *Repository) GetProject(ctx context.Context, db entrypoint.Database, projectID int) (model.Project, error) {
	var query = `
		SELECT 
		    id, 
		    creator_id,
		    description, 
		    task_config, 
		    target_overlap, 
		    tasks_per_user, 
		    created_at,
		    annotator_time_limit
		FROM project
		WHERE id = $1`
	var project ProjectTable
	err := pgxscan.Get(ctx, db, &project, query, projectID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Project{}, &storage_errors.ErrEntityNotFound{Entity: storage_errors.EntityProject}
		}
		return model.Project{}, err
	}
	return mapProjectTableToModel(project), nil
}
