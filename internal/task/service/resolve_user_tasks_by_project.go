package task_service

import (
	"context"
	ucerrors "github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	"github.com/Flak34/crowd-api/internal/pgqueue"
	project_model "github.com/Flak34/crowd-api/internal/project/model"
	projectrepo "github.com/Flak34/crowd-api/internal/project/repository"
	model "github.com/Flak34/crowd-api/internal/task/model"
	taskrepo "github.com/Flak34/crowd-api/internal/task/repository"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/riverqueue/river"
	"github.com/samber/lo"
	"time"
)

func (s *Service) ResolveUserTasksByProject(ctx context.Context, projectID int, userID int) (ResolveTasksByProjectResp, error) {
	db := s.ep.GetDB()
	resp := ResolveTasksByProjectResp{}
	project, err := s.getProject(ctx, db, projectID)
	if err != nil {
		return resp, err
	}
	reservedTasks, err := s.taskRepo.ListUserProjectTasks(ctx, db, projectID, userID)
	if err != nil {
		return resp, errors.Wrapf(ucerrors.ErrInternal, "List user project tasks: %s", err.Error())
	}
	if len(reservedTasks) != 0 {
		resp.ReservedTasks = reservedTasks
		var projectAnnotator project_model.ProjectAnnotator
		projectAnnotator, err = s.getProjectAnnotator(ctx, db, projectID, userID)
		if err != nil {
			return resp, err
		}
		resp.AnnotationDeadline = projectAnnotator.CreatedAt.UTC().Add(project.AnnotatorTimeLimit).UTC()
		return resp, nil
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
		err = s.projectRepo.CreateProjectAnnotator(ctx, tx, projectrepo.CreateProjectAnnotatorDTO{
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
		return resp, err
	}

	resp.ReservedTasks = reservedTasks
	resp.AnnotationDeadline = time.Now().UTC().Add(project.AnnotatorTimeLimit).UTC()
	return resp, nil
}
