package crowdapiv1

import (
	"context"
	crowd_api_v1 "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	projectmodel "github.com/Flak34/crowd-api/internal/project/model"
	task_model "github.com/Flak34/crowd-api/internal/task/model"
)

type TaskService interface {
	ResolveUserTasksByProject(ctx context.Context, projectID int, userID int) ([]task_model.Task, error)
}

type ProjectService interface {
	CreateProject(ctx context.Context, project projectmodel.Project) (int, error)
}

type Implementation struct {
	crowd_api_v1.UnimplementedCrowdAPIV1Server
	taskService    TaskService
	projectService ProjectService
}

func NewCrowdAPIV1(taskService TaskService, projectService ProjectService) *Implementation {
	return &Implementation{
		taskService:    taskService,
		projectService: projectService,
	}
}
