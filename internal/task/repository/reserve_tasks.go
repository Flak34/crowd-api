package task_repository

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	model "github.com/Flak34/crowd-api/internal/task/model"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type ReserveTasksDTO struct {
	UserID    int
	ProjectID int
	Limit     int
}

func (r *Repository) ReserveTasks(
	ctx context.Context,
	db entrypoint.Database,
	dto ReserveTasksDTO,
) ([]model.Task, error) {
	var query = `
		UPDATE task SET active_annotators_ids = array_append(active_annotators_ids, $1) 
        WHERE id = ANY (
            SELECT id FROM task 
            WHERE COALESCE(array_length(active_annotators_ids, 1), 0) + current_overlap < target_overlap AND 
                  project_id = $2 AND
                  deleted_at IS NULL
            LIMIT $3 
            FOR UPDATE SKIP LOCKED)
        RETURNING 
            id, 
            project_id, 
            target_overlap, 
            current_overlap, 
            active_annotators_ids, 
            input_data, 
            output_data,
            created_at;`
	var tasks []*TaskTable
	err := pgxscan.Select(ctx, db, &tasks, query, dto.UserID, dto.ProjectID, dto.Limit)
	if err != nil {
		return nil, err
	}
	res := make([]model.Task, 0, len(tasks))
	for _, task := range tasks {
		res = append(res, mapTaskTableToModel(task))
	}
	return res, nil
}
