package project_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/Flak34/crowd-api/internal/errors/storage_errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

const (
	ProjectAnnotatorPkeyName = "project_annotator_pkey"
)

type InsertProjectAnnotatorDTO struct {
	ProjectID   int
	AnnotatorID int
	TaskIDs     []int
}

func (r *Repository) InsertProjectAnnotator(
	ctx context.Context,
	db entrypoint.Database,
	dto InsertProjectAnnotatorDTO,
) error {
	var query = `
		INSERT INTO project_annotator (project_id, annotator_id, task_ids)
		VALUES 
		($1, $2, $3)`
	rows, _ := db.Query(ctx, query, dto.ProjectID, dto.AnnotatorID, dto.TaskIDs)
	rows.Close()
	err := rows.Err()
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == ProjectAnnotatorPkeyName {
			return &storage_errors.ErrDuplicateKey{UniqueConstraint: storage_errors.ProjectAnnotatorPkey}
		}
		return err
	}
	return nil
}
