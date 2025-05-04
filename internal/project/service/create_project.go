package project_service

import (
	"context"
	uscerrors "github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	model "github.com/Flak34/crowd-api/internal/project/model"
	"github.com/pkg/errors"
)

const (
	initialStatus = "created"
)

func (s *Service) CreateProject(ctx context.Context, project model.Project) (int, error) {
	db := s.ep.GetDB()
	err := validateProject(project)
	if err != nil {
		return 0, err
	}
	project.Status = initialStatus
	id, err := s.projectRepo.CreateProject(ctx, db, project)
	if err != nil {
		return 0, errors.Wrapf(uscerrors.ErrInternal, "create project: %v", err)
	}
	return id, nil
}

func validateProject(project model.Project) error {
	if len(project.Config.OutputData) == 0 {
		return errors.Wrap(uscerrors.ErrBadRequest, "Must be at least one output parameter")
	}
	for _, data := range project.Config.InputData {
		if !data.Type.IsValid() {
			return errors.Wrapf(uscerrors.ErrBadRequest, "Invalid input data type: %v", data.Type)
		}
	}
	for _, data := range project.Config.OutputData {
		if !data.Type.IsValid() {
			return errors.Wrapf(uscerrors.ErrBadRequest, "Invalid output data type: %v", data.Type)
		}
	}
	return nil
}
