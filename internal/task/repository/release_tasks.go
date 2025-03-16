package task_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) ReleaseTasks(
	ctx context.Context,
	db entrypoint.Database,
	userID int,
	taskIDs ...int,
) (updatedTasksIDs []int, err error) {
	var query = `
		UPDATE task SET active_annotators_ids = array_remove(active_annotators_ids, $1) 
        WHERE id = ANY ($2::INTEGER[]) AND $1 = ANY(active_annotators_ids)
        RETURNING id`
	var ids []int
	err = pgxscan.Select(ctx, db, &ids, query, userID, pgtype.FlatArray[int](taskIDs))
	if err != nil {
		return nil, err
	}
	return ids, nil
}
