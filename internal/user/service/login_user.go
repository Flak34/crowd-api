package user_service

import (
	"context"
	"github.com/Flak34/crowd-api/internal/errors/storage_errors"
	"github.com/Flak34/crowd-api/internal/errors/usecase_errors"
	"github.com/Flak34/crowd-api/internal/lib/jwt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) LoginUser(ctx context.Context, email, password string) (string, error) {
	db := s.ep.GetDB()
	user, err := s.repo.GetUser(ctx, db, email)
	if err != nil {
		if storage_errors.IsNotFound(err) {
			return "", errors.Wrap(usecase_errors.ErrNotFound, "Invalid credentials")
		}
		return "", err
	}
	err = bcrypt.CompareHashAndPassword(user.PassHash, []byte(password))
	if err != nil {
		return "", errors.Wrap(usecase_errors.ErrBadRequest, "Invalid credentials")
	}
	userRoles, err := s.repo.ListUserRoles(ctx, db, user.ID)
	if err != nil {
		return "", err
	}
	user.Roles = userRoles
	tokenStr, err := jwt.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
