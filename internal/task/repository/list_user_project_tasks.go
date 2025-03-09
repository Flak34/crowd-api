package task_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	model "github.com/Flak34/crowd-api/internal/task/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/samber/lo"
)

func (r *Repository) ListUserProjectTasks(
	ctx context.Context,
	db entrypoint.Database,
	projectID int,
	userID int,
) ([]model.Task, error) {
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
		FROM task JOIN (
			SELECT
				unnest(task_ids) AS task_id
			FROM project_annotator
			WHERE project_id = $1 AND annotator_id = $2) sub_query
		ON task.id = sub_query.task_id;`
	var tasks []*TaskTable
	err := pgxscan.Select(ctx, db, &tasks, query, projectID, userID)
	if err != nil {
		return nil, err
	}
	return lo.Map(tasks, func(task *TaskTable, _ int) model.Task {
		return mapTaskTableToModel(task)
	}), nil
}
