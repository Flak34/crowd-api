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

func (r *Repository) GetProjectAnnotator(
	ctx context.Context,
	db entrypoint.Database,
	projectID int,
	userID int,
) (model.ProjectAnnotator, error) {
	var query = `
		SELECT 
		    project_id, 
		    annotator_id, 
		    task_ids,
		    created_at
		FROM project_annotator
		WHERE project_id = $1 AND annotator_id = $2`
	var projectAnnotator ProjectAnnotatorTable
	err := pgxscan.Get(ctx, db, &projectAnnotator, query, projectID, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.ProjectAnnotator{}, &storage_errors.ErrEntityNotFound{Entity: storage_errors.EntityProjectAnnotator}
		}
		return model.ProjectAnnotator{}, err
	}
	return mapProjectAnnotatorToModel(projectAnnotator), nil
}
