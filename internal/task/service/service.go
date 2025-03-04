package task_service

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	projectmodel "github.com/Flak34/crowd-api/internal/project/model"
	projectrepo "github.com/Flak34/crowd-api/internal/project/repository"
	model "github.com/Flak34/crowd-api/internal/task/model"
	taskrepo "github.com/Flak34/crowd-api/internal/task/repository"
)

type TaskRepo interface {
	ReserveTasks(ctx context.Context, db entrypoint.Database, dto taskrepo.ReserveTasksDTO) ([]model.Task, error)
	ListTasks(ctx context.Context, db entrypoint.Database, taskIDs ...int) ([]model.Task, error)
}

type ProjectRepo interface {
	GetProject(ctx context.Context, db entrypoint.Database, projectID int) (projectmodel.Project, error)
	InsertProjectAnnotator(ctx context.Context, db entrypoint.Database, dto projectrepo.InsertProjectAnnotatorDTO) error
	UpdateProjectAnnotator(ctx context.Context, db entrypoint.Database, dto projectrepo.UpdateProjectAnnotatorDTO) error
	GetProjectAnnotator(ctx context.Context, db entrypoint.Database, dto projectrepo.GetProjectAnnotatorDTO) (projectmodel.ProjectAnnotator, error)
}

type Service struct {
	taskRepo    TaskRepo
	projectRepo ProjectRepo
	ep          entrypoint.Entrypoint
}

func New(ep entrypoint.Entrypoint, taskRepo TaskRepo, projectRepo ProjectRepo) *Service {
	return &Service{
		taskRepo:    taskRepo,
		projectRepo: projectRepo,
		ep:          ep,
	}
}
