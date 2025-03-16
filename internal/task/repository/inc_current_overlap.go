package task_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) IncCurrentOverlap(ctx context.Context, db entrypoint.Database, taskIDs ...int) error {
	var query = `
		UPDATE task SET current_overlap = current_overlap + 1 
        WHERE id = ANY ($1::INTEGER[])`
	rows, _ := db.Query(ctx, query, pgtype.FlatArray[int](taskIDs))
	rows.Close()
	err := rows.Err()
	if err != nil {
		return err
	}
	return nil
}
