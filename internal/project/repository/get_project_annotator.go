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

type GetProjectAnnotatorDTO struct {
	ProjectID   int
	AnnotatorID int
}

func (r *Repository) GetProjectAnnotator(
	ctx context.Context,
	db entrypoint.Database,
	dto GetProjectAnnotatorDTO,
) (model.ProjectAnnotator, error) {
	var query = `
		SELECT 
		    project_id, 
		    annotator_id, 
		    task_ids
		FROM project_annotator
		WHERE project_id = $1 AND annotator_id = $2`
	var projectAnnotator ProjectAnnotatorTable
	err := pgxscan.Get(ctx, db, &projectAnnotator, query, dto.ProjectID, dto.AnnotatorID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.ProjectAnnotator{}, &storage_errors.ErrEntityNotFound{Entity: storage_errors.EntityProjectAnnotator}
		}
		return model.ProjectAnnotator{}, err
	}
	return mapProjectAnnotatorToModel(projectAnnotator), nil
}
