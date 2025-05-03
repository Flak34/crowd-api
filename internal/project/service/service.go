package project_service

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/Flak34/crowd-api/internal/lib/common"
	model "github.com/Flak34/crowd-api/internal/project/model"
)

type projectRepo interface {
	CreateProject(ctx context.Context, db entrypoint.Database, dto model.Project) (int, error)
	ListProjects(ctx context.Context, db entrypoint.Database, opts ...common.SelectOption) ([]model.Project, error)
	CountProjects(ctx context.Context, db entrypoint.Database, opts ...common.SelectOption) (int, error)
}

type Service struct {
	ep          entrypoint.Entrypoint
	projectRepo projectRepo
}

func New(ep entrypoint.Entrypoint, projectRepo projectRepo) *Service {
	return &Service{
		ep:          ep,
		projectRepo: projectRepo,
	}
}
