package crowdapiv1

import (
	"context"
	crowd_api_v1 "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	task_model "github.com/Flak34/crowd-api/internal/task/model"
)

type TaskService interface {
	ResolveUserTasksByProject(ctx context.Context, projectID int, userID int) ([]task_model.Task, error)
}

type Implementation struct {
	crowd_api_v1.UnimplementedCrowdAPIV1Server
	taskService TaskService
}

func NewCrowdAPIV1(taskService TaskService) *Implementation {
	return &Implementation{
		taskService: taskService,
	}
}
