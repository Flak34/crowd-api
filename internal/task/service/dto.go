package task_service

import (
	model "github.com/Flak34/crowd-api/internal/task/model"
	"time"
)

type ResolveTasksByProjectResp struct {
	ReservedTasks      []model.Task
	AnnotationDeadline time.Time
}
