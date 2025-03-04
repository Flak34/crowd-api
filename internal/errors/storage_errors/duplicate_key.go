package storage_errors

import "fmt"

type UniqueConstraintName string

const (
	ProjectAnnotatorPkey      UniqueConstraintName = "project_annotator_pkey"
	TaskAnnotatorUqConstraint UniqueConstraintName = "task_annotation_pkey"
)

type ErrDuplicateKey struct {
	UniqueConstraint UniqueConstraintName
}

func (e *ErrDuplicateKey) Error() string {
	return fmt.Sprintf("unique constraint %s violation", e.UniqueConstraint)
}
