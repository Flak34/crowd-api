package annotation_service

import (
	"context"
	model "github.com/Flak34/crowd-api/internal/annotation/model"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	projectmodel "github.com/Flak34/crowd-api/internal/project/model"
	taskmodel "github.com/Flak34/crowd-api/internal/task/model"
)

type taskRepository interface {
	ReleaseTasks(ctx context.Context, db entrypoint.Database, userID int, taskIDs ...int) ([]int, error)
	IncCurrentOverlap(ctx context.Context, db entrypoint.Database, taskIDs ...int) error
	ListTasks(ctx context.Context, db entrypoint.Database, taskIDs ...int) ([]taskmodel.Task, error)
}

type annotationRepository interface {
	CreateAnnotations(ctx context.Context, db entrypoint.Database, annotations ...model.Annotation) error
}

type projectRepository interface {
	GetProject(ctx context.Context, db entrypoint.Database, projectID int) (projectmodel.Project, error)
	GetProjectAnnotator(ctx context.Context, db entrypoint.Database, projectID int, userID int) (projectmodel.ProjectAnnotator, error)
}

type Service struct {
	ep entrypoint.Entrypoint

	taskRepo       taskRepository
	annotationRepo annotationRepository
	projectRepo    projectRepository
}

func New(
	ep entrypoint.Entrypoint,
	taskRepo taskRepository,
	annotationRepo annotationRepository,
	projectRepo projectRepository,
) *Service {
	return &Service{
		ep:             ep,
		taskRepo:       taskRepo,
		annotationRepo: annotationRepo,
		projectRepo:    projectRepo,
	}
}
