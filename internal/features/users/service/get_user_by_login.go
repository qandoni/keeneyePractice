package users_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *UsersService) GetUserByLogin(
	ctx context.Context,
	login string,
) (domain.User, error) {
	user, err := s.usersRepository.GetUserByLogin(ctx, login)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}
	return user, nil
}
