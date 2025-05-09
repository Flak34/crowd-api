package crowdapiv1

import (
	"context"
	crowdapiv1 "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	projectmodel "github.com/Flak34/crowd-api/internal/project/model"
	projectservice "github.com/Flak34/crowd-api/internal/project/service"
	task_service "github.com/Flak34/crowd-api/internal/task/service"
	"io"
)

type taskService interface {
	ResolveUserTasksByProject(ctx context.Context, projectID int, userID int) (task_service.ResolveTasksByProjectResp, error)
}

type projectService interface {
	CreateProject(ctx context.Context, project projectmodel.Project) (int, error)
	ListProjects(ctx context.Context, dto projectservice.ListProjectsDTO) (projects []projectmodel.Project, totalCount int, err error)
	GetProject(ctx context.Context, projectID int) (projectmodel.Project, error)
	UploadProjectData(ctx context.Context, projectID int, dataReader io.Reader) error
}

type Implementation struct {
	crowdapiv1.UnimplementedCrowdAPIV1Server
	taskService    taskService
	projectService projectService
}

func NewCrowdAPIV1(taskService taskService, projectService projectService) *Implementation {
	return &Implementation{
		taskService:    taskService,
		projectService: projectService,
	}
}
