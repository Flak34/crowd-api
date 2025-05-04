package project_service

import (
	"context"
	"github.com/Flak34/crowd-api/internal/errors/storage_errors"
	"github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	model "github.com/Flak34/crowd-api/internal/project/model"
	"github.com/pkg/errors"
)

func (s *Service) GetProject(ctx context.Context, projectID int) (model.Project, error) {
	db := s.ep.GetDB()
	project, err := s.projectRepo.GetProject(ctx, db, projectID)
	if err != nil {
		if storage_errors.IsNotFound(err) {
			return project, errors.Wrapf(usecase_errors.ErrNotFound, "get project by id: %d", projectID)
		}
		return project, err
	}
	return project, nil
}
