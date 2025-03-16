package task_service

import (
	"context"
	"github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	"github.com/Flak34/crowd-api/internal/pgqueue"
	"github.com/pkg/errors"
	"github.com/riverqueue/river"
	"github.com/rs/zerolog/log"
	"time"
)

func (s *Service) ReleaseUserProjectTasks(ctx context.Context, projectID int, userID int) error {
	db := s.ep.GetDB()
	project, err := s.getProject(ctx, db, projectID)
	if err != nil {
		return err
	}
	projectAnnotator, err := s.getProjectAnnotator(ctx, db, projectID, userID)
	if err != nil {
		return err
	}

	annotationDeadline := projectAnnotator.CreatedAt.Add(project.AnnotatorTimeLimit)
	timeNow := time.Now()
	if timeNow.Before(annotationDeadline) {
		log.Warn().
			Int("ProjectID", projectID).
			Int("UserID", userID).
			Time("AnnotationDeadline", annotationDeadline).
			Time("CurrentTime", timeNow).
			Msg("Tasks cannot be released before annotation deadline")
		_, err = s.pgqClient.Insert(ctx, pgqueue.AnnotationDeadlineArgs{
			ProjectID:   projectID,
			AnnotatorID: userID,
		}, &river.InsertOpts{
			ScheduledAt: annotationDeadline,
		})
		return err
	}

	_, err = s.taskRepo.ReleaseTasks(ctx, db, userID, projectAnnotator.TaskIDs...)
	if err != nil {
		return errors.Wrapf(usecase_errors.ErrInternal, "release user tasks: %s", err.Error())
	}
	return nil
}
