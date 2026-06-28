package auth_service

import (
	"context"
	"fmt"

	auth_contracts "github.com/qandoni/keeneyePractice/internal/features/auth/contracts"
)

func (s *AuthService) Login(
	ctx context.Context,
	input auth_contracts.LoginInput,
) (string, error) {
	user, err := s.usersRepository.GetUserByLogin(
		ctx,
		input.Login,
	)
	if err != nil {
		return "", fmt.Errorf("get user by login: %w", err)
	}

	err = s.passwordHasher.Compare(
		user.PasswordHash,
		input.Password,
	)

	if err != nil {
		return "", fmt.Errorf("compare password: %w", err)
	}

	token, err := s.jwtManager.GenerateToken(user)
	if err != nil {
		return "", fmt.Errorf("generate JWT token: %w", err)
	}
	return token, nil
}
