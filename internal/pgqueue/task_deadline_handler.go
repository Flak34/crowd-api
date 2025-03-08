package pgqueue

import (
	"context"
	"github.com/riverqueue/river"
	"github.com/rs/zerolog/log"
)

const (
	TaskDeadlineArgsKind = "task_deadline"
)

type AnnotationDeadlineArgs struct {
	ProjectID   int
	AnnotatorID int
}

func (AnnotationDeadlineArgs) Kind() string { return TaskDeadlineArgsKind }

type AnnotationDeadlineHandler struct {
	river.WorkerDefaults[AnnotationDeadlineArgs]
}

func (h *AnnotationDeadlineHandler) Work(ctx context.Context, job *river.Job[AnnotationDeadlineArgs]) error {
	// 1. Достаю ProjectAnnotator
	// 2. Достаю Project
	// 3. Достаю все таски пользака
	// 4. Если затаймаутился, то удаляю пользака из всех его задач
	log.Info().Msgf("process AnnotationDeadlineHandler: %v", job.Args)
	return nil
}
