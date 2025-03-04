package usecase_errors

import "github.com/pkg/errors"

var ErrNotFound = errors.New("not found")

var ErrBadRequest = errors.New("bad request")

var ErrUnauthorized = errors.New("unauthorized")

var ErrForbidden = errors.New("forbidden")

var ErrInternal = errors.New("internal error")
