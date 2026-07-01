package auth_service

import (
	"context"
	"fmt"
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	auth_contracts "github.com/qandoni/keeneyePractice/internal/features/auth/contracts"
)

func (s *AuthService) Login(
	ctx context.Context,
	input auth_contracts.LoginInput,
) (auth_contracts.LoginOutput, error) {
	user, err := s.usersRepository.GetUserByLogin(
		ctx,
		input.Login,
	)
	if err != nil {
		return auth_contracts.LoginOutput{}, fmt.Errorf("get user by login: %w", err)
	}

	err = s.passwordHasher.Compare(
		user.PasswordHash,
		input.Password,
	)

	if err != nil {
		return auth_contracts.LoginOutput{}, fmt.Errorf("compare password: %w", err)
	}

	accessToken, err := s.jwtManager.GenerateAccessToken(user)
	if err != nil {
		return auth_contracts.LoginOutput{}, err
	}

	refreshToken, err := s.refreshGenerator.Generate()
	if err != nil {
		return auth_contracts.LoginOutput{}, err
	}

	refreshHash := s.sha256Hasher.Hash(refreshToken)

	err = s.refreshRepository.Create(
		ctx,
		domain.RefreshToken{
			TokenHash: refreshHash,
			UserID:    user.ID,
			ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
		},
	)
	if err != nil {
		return auth_contracts.LoginOutput{}, err
	}

	return auth_contracts.LoginOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
