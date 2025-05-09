package project_service

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/Flak34/crowd-api/internal/lib/common"
	model "github.com/Flak34/crowd-api/internal/project/model"
	task_model "github.com/Flak34/crowd-api/internal/task/model"
)

type projectRepo interface {
	CreateProject(ctx context.Context, db entrypoint.Database, dto model.Project) (int, error)
	ListProjects(ctx context.Context, db entrypoint.Database, opts ...common.SelectOption) ([]model.Project, error)
	CountProjects(ctx context.Context, db entrypoint.Database, opts ...common.SelectOption) (int, error)
	GetProject(ctx context.Context, db entrypoint.Database, projectID int) (model.Project, error)
	UpdateStatus(ctx context.Context, db entrypoint.Database, projectID int, newStatus string) error
}

type taskRepo interface {
	CreateTasks(ctx context.Context, db entrypoint.Database, tasks ...task_model.Task) error
}

type Service struct {
	ep          entrypoint.Entrypoint
	projectRepo projectRepo
	taskRepo    taskRepo
}

func New(ep entrypoint.Entrypoint, projectRepo projectRepo, taskRepo taskRepo) *Service {
	return &Service{
		ep:          ep,
		projectRepo: projectRepo,
		taskRepo:    taskRepo,
	}
}
