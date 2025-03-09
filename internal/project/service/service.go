package project_service

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	projectrepo "github.com/Flak34/crowd-api/internal/project/repository"
)

type ProjectRepo interface {
	CreateProject(ctx context.Context, db entrypoint.Database, dto projectrepo.CreateProjectDTO) (int, error)
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
