package project_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/jackc/pgx/v5/pgtype"
)

type UpdateProjectAnnotatorDTO struct {
	ProjectID   int
	AnnotatorID int
	TaskIDs     []int
}

func (r *Repository) UpdateProjectAnnotator(
	ctx context.Context,
	db entrypoint.Database,
	dto UpdateProjectAnnotatorDTO,
) error {
	var query = `UPDATE project_annotator SET task_ids = $1 WHERE project_id = $2 AND annotator_id = $3`
	rows, err := db.Query(ctx, query, pgtype.FlatArray[int](dto.TaskIDs), dto.ProjectID, dto.AnnotatorID)
	rows.Close()
	if err != nil {
		return err
	}
	return nil
}
