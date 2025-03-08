package task_service

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	dberrors "github.com/Flak34/crowd-api/internal/errors/storage_errors"
	ucerrors "github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	"github.com/Flak34/crowd-api/internal/pgqueue"
	projectmodel "github.com/Flak34/crowd-api/internal/project/model"
	projectrepo "github.com/Flak34/crowd-api/internal/project/repository"
	model "github.com/Flak34/crowd-api/internal/task/model"
	taskrepo "github.com/Flak34/crowd-api/internal/task/repository"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/riverqueue/river"
	"github.com/samber/lo"
	"time"
)

func (s *Service) ResolveUserTasksByProject(ctx context.Context, projectID int, userID int) ([]model.Task, error) {
	db := s.ep.GetDB()
	project, err := s.getProject(ctx, db, projectID)
	if err != nil {
		return nil, err
	}
	reservedTasks, err := s.listUserTasks(ctx, db, userID, projectID)
	if err != nil {
		return nil, err
	}
	if len(reservedTasks) != 0 {
		return reservedTasks, nil
	}
	err = s.ep.TxWrapper(ctx, func(ctx context.Context, tx pgx.Tx) error {
		reservedTasks, err = s.taskRepo.ReserveTasks(ctx, tx, taskrepo.ReserveTasksDTO{
			UserID:    userID,
			ProjectID: projectID,
			Limit:     project.TasksPerUser,
		})
		if err != nil {
			return err
		}
		if len(reservedTasks) == 0 {
			return nil
		}
		err = s.projectRepo.InsertProjectAnnotator(ctx, tx, projectrepo.InsertProjectAnnotatorDTO{
			ProjectID:   projectID,
			AnnotatorID: userID,
			TaskIDs: lo.Map(reservedTasks, func(task model.Task, _ int) int {
				return task.ID
			}),
		})
		if err != nil {
			return err
		}
		_, err = s.pgqClient.InsertTx(ctx, tx, pgqueue.AnnotationDeadlineArgs{
			ProjectID:   projectID,
			AnnotatorID: userID,
		}, &river.InsertOpts{ScheduledAt: time.Now().Add(project.AnnotatorTimeLimit)})
		return err
	})
	if err != nil {
		return reservedTasks, err
	}

	return reservedTasks, nil
}

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

func (s *Service) listUserTasks(
	ctx context.Context,
	db entrypoint.Database,
	userID int,
	projectID int,
) ([]model.Task, error) {
	projectAnnotator, err := s.projectRepo.GetProjectAnnotator(ctx, db, projectrepo.GetProjectAnnotatorDTO{
		ProjectID:   projectID,
		AnnotatorID: userID,
	})
	if err != nil {
		if dberrors.IsNotFound(err) {
			return []model.Task{}, nil
		}
		return []model.Task{}, errors.Wrapf(ucerrors.ErrInternal, "Get project annotator: %s", err.Error())
	}
	tasks, err := s.taskRepo.ListTasks(ctx, db, projectAnnotator.TaskIDs...)
	if err != nil {
		return []model.Task{}, errors.Wrapf(ucerrors.ErrInternal, "List tasks: %s", err.Error())
	}
	return tasks, nil
}
