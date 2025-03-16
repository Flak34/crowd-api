package annotation_service

import (
	"context"
	model "github.com/Flak34/crowd-api/internal/annotation/model"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	dberrors "github.com/Flak34/crowd-api/internal/errors/storage_errors"
	uscerrors "github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	taskmodel "github.com/Flak34/crowd-api/internal/task/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"maps"
	"slices"
	"time"
)

func (s *Service) CreateAnnotations(ctx context.Context, annotations ...model.Annotation) error {
	// Создание аннотаций
	// Увеличнение текущего оверлапа задач
	// Освобождение этих задач, те удаление пользователя из массива active_annotators_ids
	db := s.ep.GetDB()
	taskIDs := lo.Map(annotations, func(annotation model.Annotation, _ int) int {
		return annotation.TaskID
	})
	taskByID, err := s.getTasksMap(ctx, db, taskIDs...)
	if err != nil {
		return err
	}

	projectIDs := lo.Map(slices.Collect(maps.Values(taskByID)), func(task taskmodel.Task, _ int) int {
		return task.ProjectID
	})
	if len(projectIDs) > 1 {
		return errors.Wrapf(uscerrors.ErrBadRequest,
			"It is prohibited to annotate tasks from different projects in one submission")
	}
	project, err := s.projectRepo.GetProject(ctx, db, projectIDs[0])
	if err != nil {
		if dberrors.IsNotFound(err) {
			return errors.Wrapf(uscerrors.ErrNotFound, "Can't find project by id: %d", projectIDs[0])
		}
		return errors.Wrapf(uscerrors.ErrInternal, "get project: %s", err.Error())
	}

	projectAnnotator, err := s.projectRepo.GetProjectAnnotator(ctx, db, projectIDs[0], annotations[0].AnnotatorID)
	if err != nil {
		if dberrors.IsNotFound(err) {
			return errors.Wrapf(uscerrors.ErrNotFound,
				"Can't find project annotator (userID: %d projectID: %d)",
				annotations[0].AnnotatorID, projectIDs[0])
		}
		return errors.Wrapf(uscerrors.ErrInternal, "get project: %s", err.Error())
	}

	timeNow := time.Now()
	for _, annotation := range annotations {
		task, taskExists := taskByID[annotation.TaskID]
		if !taskExists {
			return errors.Wrap(uscerrors.ErrBadRequest, "Annotation can be created only for active task")
		}
		validationErr := project.ValidateAnnotation(annotation)
		if validationErr != nil {
			return errors.Wrapf(uscerrors.ErrBadRequest, "Validate annotation: %s", validationErr.Error())
		}
	}

	return nil
}

func (s *Service) getTasksMap(
	ctx context.Context,
	db entrypoint.Database,
	taskIDs ...int,
) (map[int]taskmodel.Task, error) {
	tasks, err := s.taskRepo.ListTasks(ctx, db, taskIDs...)
	if err != nil {
		return nil, errors.Wrapf(uscerrors.ErrInternal, "list tasks: %s", err.Error())
	}
	taskByID := make(map[int]taskmodel.Task, len(tasks))
	for _, task := range tasks {
		taskByID[task.ID] = task
	}
	return taskByID, nil
}

func (s *Service) createAnnotations(ctx context.Context, annotations []model.Annotation, taskIDs int) error {
	err := s.ep.TxWrapper(ctx, func(ctx context.Context, tx pgx.Tx) error {

	})
	if err != nil {
		return err
	}
	return nil
}
