package storage_errors

import (
	"fmt"
	"github.com/pkg/errors"
)

type EntityName string

const (
	EntityProject          EntityName = "project"
	EntityProjectAnnotator EntityName = "project_annotator"
	EntityTask             EntityName = "task"
	EntityUser             EntityName = "user"
)

type ErrEntityNotFound struct {
	Entity EntityName
}

func (e *ErrEntityNotFound) Error() string {
	return fmt.Sprintf("entity %s not found", e.Entity)
}

func IsNotFound(err error) bool {
	var e *ErrEntityNotFound
	return errors.As(err, &e)
}
