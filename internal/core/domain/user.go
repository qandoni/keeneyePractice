package domain

import (
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

type User struct {
	ID           int
	Version      int
	Login        string
	PasswordHash string
	Role         enum.Role
}

func NewUser(
	id int,
	version int,
	login string,
	passwordHash string,
	role string,
) User {
	return User{
		ID:           id,
		Version:      version,
		Login:        login,
		PasswordHash: passwordHash,
		Role:         enum.Role(role),
	}
}

func (u *User) Validate() error {
	loginLen := len([]rune(u.Login))
	if loginLen < 3 || loginLen > 100 {
		return fmt.Errorf("invalid `Login` len: %d: %w", loginLen, core_errors.ErrInvalidArgument)
	}
	return nil
}

func NewUserUnitialized(
	login string,
	passwordHash string,
	role string,
) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		login,
		passwordHash,
		role,
	)
}

type UserPatch struct {
	Login    Nullable[string]
	Password Nullable[string]
	Role     Nullable[enum.Role]
}

func NewUserPatch(
	login Nullable[string],
	password Nullable[string],
	role Nullable[enum.Role],
) UserPatch {
	return UserPatch{
		Login:    login,
		Password: password,
		Role:     role,
	}
}

func (p *UserPatch) Validate() error {
	if p.Login.Set && p.Login.Value == nil {
		return fmt.Errorf("`Login` can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	if p.Password.Set && p.Password.Value == nil {
		return fmt.Errorf("'Password' can't be patched to NULL : %w", core_errors.ErrInvalidArgument)
	}

	if p.Role.Set && p.Role.Value == nil {
		return fmt.Errorf("'Role' can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}
