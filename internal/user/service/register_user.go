package user_service

import (
	"context"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserDTO struct {
	Email     string
	Password  string
	RoleNames []string
}

func (s *Service) RegisterUser(ctx context.Context, dto CreateUserDTO) error {
	err := s.ep.TxWrapper(ctx, func(ctx context.Context, tx pgx.Tx) error {
		roles, err := s.repo.ListRolesByName(ctx, tx, dto.RoleNames...)
		if err != nil {
			return err
		}
		passHash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		userID, err := s.repo.CreateUser(ctx, tx, dto.Email, passHash)
		if err != nil {
			return err
		}
		err = s.repo.InsertUserRoles(ctx, tx, userID, roles)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
