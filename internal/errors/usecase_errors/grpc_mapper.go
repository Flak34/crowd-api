package usecase_errors

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapToGRPC(err error) error {
	switch {
	case errors.Is(err, ErrNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, ErrForbidden):
		return status.Error(codes.PermissionDenied, err.Error())
	case errors.Is(err, ErrUnauthorized):
		return status.Error(codes.Unauthenticated, err.Error())
	case errors.Is(err, ErrBadRequest):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrInternal):
		return status.Error(codes.Internal, err.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}
