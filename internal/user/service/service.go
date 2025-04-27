package user_service

import (
	"context"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	"github.com/Flak34/crowd-api/internal/user/model"
)

type userRepo interface {
	GetUser(ctx context.Context, db entrypoint.Database, email string) (user_model.User, error)
	CreateUser(ctx context.Context, db entrypoint.Database, email string, passHash []byte) (int, error)
	ListRolesByName(ctx context.Context, db entrypoint.Database, roleNames ...string) ([]user_model.Role, error)
	ListUserRoles(ctx context.Context, db entrypoint.Database, userID int) ([]user_model.Role, error)
	InsertUserRoles(ctx context.Context, db entrypoint.Database, userID int, roles []user_model.Role) error
}

type Service struct {
	repo userRepo
	ep   entrypoint.Entrypoint
}

func New(ep entrypoint.Entrypoint, repo userRepo) *Service {
	return &Service{repo: repo, ep: ep}
}
