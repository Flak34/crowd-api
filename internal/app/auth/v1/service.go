package authv1

import (
	"context"
	auth_v1 "github.com/Flak34/crowd-api/internal/pb/auth"
	"github.com/Flak34/crowd-api/internal/user/service"
)

type userService interface {
	RegisterUser(ctx context.Context, dto user_service.CreateUserDTO) error
	LoginUser(ctx context.Context, email, password string) (string, error)
}

type Implementation struct {
	auth_v1.UnimplementedAuthV1Server

	userService userService
}

func NewAuthV1(service userService) *Implementation {
	return &Implementation{
		userService: service,
	}
}
