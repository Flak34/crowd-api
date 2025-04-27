package authv1

import (
	"context"
	"github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	auth_v1 "github.com/Flak34/crowd-api/internal/pb/auth"
)

func (i *Implementation) Login(ctx context.Context, req *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	tokenStr, err := i.userService.LoginUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, usecase_errors.MapToGRPC(err)
	}
	return &auth_v1.LoginResponse{Token: tokenStr}, nil
}
