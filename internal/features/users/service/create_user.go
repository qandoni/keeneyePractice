package users_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	users_contracts "github.com/qandoni/keeneyePractice/internal/features/users/contracts"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	input users_contracts.CreateUserInput,
) (domain.User, error) {
	if err := input.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate input data: %w", err)
	}
	passwordHash, err := s.passwordHasher.Hash(input.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash password: %w", err)
	}

	user := domain.NewUserUnitialized(input.Login, passwordHash, string(input.Role))

	user, err = s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user in repository: %w", err)
	}

	return user, nil
}
