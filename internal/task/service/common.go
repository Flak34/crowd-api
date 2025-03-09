package task_service

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	dberrors "github.com/Flak34/crowd-api/internal/errors/storage_errors"
	ucerrors "github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	projectmodel "github.com/Flak34/crowd-api/internal/project/model"
	"github.com/pkg/errors"
)

func (s *Service) getProject(
	ctx context.Context,
	db entrypoint.Database,
	projectID int,
) (projectmodel.Project, error) {
	project, err := s.projectRepo.GetProject(ctx, db, projectID)
	if err != nil {
		if dberrors.IsNotFound(err) {
			return project, errors.Wrapf(ucerrors.ErrNotFound, "Get project")
		}
		return project, errors.Wrapf(ucerrors.ErrInternal, "Get project: %s", err.Error())
	}
	return project, nil
}

func (s *Service) getProjectAnnotator(
	ctx context.Context,
	db entrypoint.Database,
	projectID int,
	userID int,
) (projectmodel.ProjectAnnotator, error) {
	projectAnnotator, err := s.projectRepo.GetProjectAnnotator(ctx, db, projectID, userID)
	if err != nil {
		if dberrors.IsNotFound(err) {
			return projectAnnotator, errors.Wrapf(ucerrors.ErrNotFound, "Get project annotator")
		}
		return projectAnnotator, errors.Wrapf(ucerrors.ErrInternal, "Get project annotator: %s", err.Error())
	}
	return projectAnnotator, nil
}
