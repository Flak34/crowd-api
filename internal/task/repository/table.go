package task_repository

import (
	"database/sql"
	model "github.com/Flak34/crowd-api/internal/task/model"
	"time"
)

type TaskTable struct {
	ID                  int            `db:"id"`
	ProjectID           int            `db:"project_id"`
	TargetOverlap       int            `db:"target_overlap"`
	CurrentOverlap      int            `db:"current_overlap"`
	ActiveAnnotatorsIDs []int          `db:"active_annotators_ids"`
	InputData           string         `db:"input_data"`
	OutputData          sql.NullString `db:"output_data"`
	CreatedAt           time.Time      `db:"created_at"`
}

func mapTaskTableToModel(task *TaskTable) model.Task {
	if task == nil {
		return model.Task{}
	}
	taskModel := model.Task{
		ID:                  task.ID,
		ProjectID:           task.ProjectID,
		TargetOverlap:       task.TargetOverlap,
		CurrentOverlap:      task.CurrentOverlap,
		ActiveAnnotatorsIDs: task.ActiveAnnotatorsIDs,
		InputData:           task.InputData,
		CreatedAt:           task.CreatedAt,
	}
	if task.OutputData.Valid {
		taskModel.OutputData = task.OutputData.String
	}
	return taskModel
}
