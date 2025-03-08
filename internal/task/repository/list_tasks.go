package task_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	model "github.com/Flak34/crowd-api/internal/task/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
)

func (r *Repository) ListTasks(ctx context.Context, db entrypoint.Database, taskIDs ...int) ([]model.Task, error) {
	var query = `
		SELECT 
		    id,
		    project_id,
		    target_overlap, 
		    current_overlap,
		    active_annotators_ids, 
		    input_data,
		    output_data,
		    created_at
		FROM task
		WHERE id = ANY($1::INTEGER[])`
	var tasks []*TaskTable
	err := pgxscan.Select(ctx, db, &tasks, query, pgtype.FlatArray[int](taskIDs))
	if err != nil {
		return nil, err
	}
	return lo.Map(tasks, func(task *TaskTable, _ int) model.Task {
		return mapTaskTableToModel(task)
	}), nil
}
