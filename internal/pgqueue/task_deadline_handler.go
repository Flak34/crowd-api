package pgqueue

import (
	"context"
	ucerrors "github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	TaskDeadlineArgsKind = "task_deadline"
)

type AnnotationDeadlineArgs struct {
	ProjectID   int
	AnnotatorID int
}

func (AnnotationDeadlineArgs) Kind() string { return TaskDeadlineArgsKind }

type taskService interface {
	ReleaseUserProjectTasks(ctx context.Context, projectID int, userID int) error
}

type AnnotationDeadlineHandler struct {
	taskService taskService
}

func NewAnnotationDeadlineHandler(taskService taskService) *AnnotationDeadlineHandler {
	return &AnnotationDeadlineHandler{taskService: taskService}
}

func (h *AnnotationDeadlineHandler) Work(ctx context.Context, job *river.Job[AnnotationDeadlineArgs]) error {
	log.Info().Int("ProjectID", job.Args.ProjectID).
		Int("AnnotatorID", job.Args.AnnotatorID).
		Msg("process AnnotationDeadlineHandler")

	err := h.taskService.ReleaseUserProjectTasks(ctx, job.Args.ProjectID, job.Args.AnnotatorID)
	if err != nil {
		log.Error().Err(err).Msg("failed to release user project tasks")
		if ucerrors.IsInternal(err) {
			return err
		}
	}
	return nil
}

func (h *AnnotationDeadlineHandler) Middleware(*river.Job[AnnotationDeadlineArgs]) []rivertype.WorkerMiddleware {
	return nil
}

func (h *AnnotationDeadlineHandler) NextRetry(*river.Job[AnnotationDeadlineArgs]) time.Time {
	return time.Time{}
}

func (h *AnnotationDeadlineHandler) Timeout(*river.Job[AnnotationDeadlineArgs]) time.Duration {
	return 0
}
