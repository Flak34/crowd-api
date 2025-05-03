package crowdapiv1

import (
	"context"
	crowdapiv1 "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	projectmodel "github.com/Flak34/crowd-api/internal/project/model"
	projectservice "github.com/Flak34/crowd-api/internal/project/service"
	taskmodel "github.com/Flak34/crowd-api/internal/task/model"
)

type TaskService interface {
	ResolveUserTasksByProject(ctx context.Context, projectID int, userID int) ([]taskmodel.Task, error)
}

type ProjectService interface {
	CreateProject(ctx context.Context, project projectmodel.Project) (int, error)
	ListProjects(ctx context.Context, dto projectservice.ListProjectsDTO) (projects []projectmodel.Project, totalCount int, err error)
}

type Implementation struct {
	crowdapiv1.UnimplementedCrowdAPIV1Server
	taskService    TaskService
	projectService ProjectService
}

func NewCrowdAPIV1(taskService TaskService, projectService ProjectService) *Implementation {
	return &Implementation{
		taskService:    taskService,
		projectService: projectService,
	}
}
