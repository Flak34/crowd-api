package annotation_repository

import (
	"context"
	model "github.com/Flak34/crowd-api/internal/annotation/model"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/Flak34/crowd-api/internal/errors/storage_errors"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

const (
	TaskAnnotationPkeyName = "task_annotation_pkey"
)

func (r *Repository) CreateAnnotations(
	ctx context.Context,
	db entrypoint.Database,
	annotations ...model.Annotation,
) error {
	insBuilder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	insBuilder.InsertInto("task_annotation").Cols(
		"task_id",
		"annotator_id",
		"output_data",
	)
	for _, annotation := range annotations {
		insBuilder.Values(
			annotation.TaskID,
			annotation.AnnotatorID,
			annotation.OutputData,
		)
	}
	q, args := insBuilder.Build()
	rows, _ := db.Query(ctx, q, args...)
	rows.Close()
	err := rows.Err()
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == TaskAnnotationPkeyName {
			return &storage_errors.ErrDuplicateKey{UniqueConstraint: storage_errors.TaskAnnotatorUqConstraint}
		}
		return err
	}
	return nil
}
