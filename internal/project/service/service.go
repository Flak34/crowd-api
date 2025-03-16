package project_service

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	model "github.com/Flak34/crowd-api/internal/project/model"
)

type ProjectRepo interface {
	CreateProject(ctx context.Context, db entrypoint.Database, dto model.Project) (int, error)
}

type Service struct {
	ep          entrypoint.Entrypoint
	projectRepo ProjectRepo
}

func New(ep entrypoint.Entrypoint, projectRepo ProjectRepo) *Service {
	return &Service{
		ep:          ep,
		projectRepo: projectRepo,
	}
}
