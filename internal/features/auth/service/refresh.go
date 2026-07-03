package auth_service

import (
	"context"
	"fmt"
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	auth_contracts "github.com/qandoni/keeneyePractice/internal/features/auth/contracts"
)

func (s *AuthService) Refresh(
	ctx context.Context,
	input auth_contracts.RefreshInput,
) (auth_contracts.RefreshOutput, error) {

	hash := s.sha256Hasher.Hash(input.RefreshToken)

	refreshToken, err := s.refreshRepository.Get(ctx, hash)
	if err != nil {
		return auth_contracts.RefreshOutput{}, fmt.Errorf(
			"get refresh token: %w",
			err,
		)
	}

	if refreshToken.ExpiresAt.Before(time.Now()) {
		err = s.txManager.WithinTransaction(ctx, func(ctx context.Context) error {
			return s.refreshRepository.Delete(ctx, hash)
		})
		if err != nil {
			return auth_contracts.RefreshOutput{}, err
		}

		return auth_contracts.RefreshOutput{}, core_errors.ErrRefreshTokenExpired
	}

	user, err := s.usersRepository.GetUser(ctx, refreshToken.UserID)
	if err != nil {
		return auth_contracts.RefreshOutput{}, fmt.Errorf(
			"get user: %w",
			err,
		)
	}

	accessToken, err := s.jwtManager.GenerateAccessToken(user)
	if err != nil {
		return auth_contracts.RefreshOutput{}, fmt.Errorf(
			"generate access token: %w",
			err,
		)
	}

	newRefreshToken, err := s.refreshGenerator.Generate()
	if err != nil {
		return auth_contracts.RefreshOutput{}, fmt.Errorf(
			"generate refresh token: %w",
			err,
		)
	}

	newHash := s.sha256Hasher.Hash(newRefreshToken)

	err = s.txManager.WithinTransaction(ctx, func(ctx context.Context) error {

		err := s.refreshRepository.Save(
			ctx,
			domain.NewRefreshToken(
				refreshToken.ID,
				refreshToken.Version,
				refreshToken.UserID,
				newHash,
				time.Now().Add(30*24*time.Hour),
			),
		)
		if err != nil {
			return fmt.Errorf(
				"save refresh token: %w",
				err,
			)
		}

		return nil
	})
	if err != nil {
		return auth_contracts.RefreshOutput{}, err
	}

	return auth_contracts.RefreshOutput{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
