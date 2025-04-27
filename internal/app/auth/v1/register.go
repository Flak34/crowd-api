package authv1

import (
	"context"
	"github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	auth_v1 "github.com/Flak34/crowd-api/internal/pb/auth"
	"github.com/Flak34/crowd-api/internal/user/service"
)

func (i *Implementation) Register(ctx context.Context, req *auth_v1.RegisterRequest) (*auth_v1.RegisterResponse, error) {
	err := i.userService.RegisterUser(ctx, user_service.CreateUserDTO{
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
		RoleNames: req.GetUserRoles(),
	})
	if err != nil {
		return nil, usecase_errors.MapToGRPC(err)
	}
	return &auth_v1.RegisterResponse{}, nil
}
