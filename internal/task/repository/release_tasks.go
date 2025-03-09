package task_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) ReleaseTasks(
	ctx context.Context,
	db entrypoint.Database,
	userID int,
	taskIDs ...int,
) error {
	var query = `
		UPDATE task SET active_annotators_ids = array_remove(active_annotators_ids, $1) 
        WHERE id = ANY ($2::INTEGER[])`
	rows, _ := db.Query(ctx, query, userID, pgtype.FlatArray[int](taskIDs))
	rows.Close()
	err := rows.Err()
	if err != nil {
		return err
	}
	return nil
}
