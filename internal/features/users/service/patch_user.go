package users_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *UsersService) PatchUser(
	ctx context.Context,
	id int,
	patch domain.UserPatch,
) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}
	if user, err = s.ApplyPatch(user, patch); err != nil {
		return domain.User{}, fmt.Errorf("apply user patch: %w", err)
	}
	patchedUser, err := s.usersRepository.PatchUser(ctx, id, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("patch user: %w", err)
	}
	return patchedUser, nil
}

func (s *UsersService) ApplyPatch(
	user domain.User,
	patch domain.UserPatch,
) (domain.User, error) {
	tmp := user
	if patch.Login.Set {
		tmp.Login = *patch.Login.Value
	}
	if patch.Role.Set {
		tmp.Role = *patch.Role.Value
	}
	if patch.Password.Set {
		if patch.Password.Value == nil {
			return domain.User{}, fmt.Errorf("password can't be NULL")
		}
		hashed, err := s.passwordHasher.Hash(*patch.Password.Value)
		if err != nil {
			return domain.User{}, fmt.Errorf("hash password: %w", err)
		}
		tmp.PasswordHash = hashed
	}

	if err := tmp.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate patched user: %w", err)
	}
	return tmp, nil
}
