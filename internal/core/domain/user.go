package domain

import (
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

type User struct {
	ID           int
	Version      int
	Email        string
	PasswordHash string
	Role         enum.Role
}

func NewUser(
	id int,
	version int,
	email string,
	passwordHash string,
	role string,
) User {
	return User{
		ID:           id,
		Version:      version,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         enum.Role(role),
	}
}

func (u *User) Validate() error {
	emailLen := len([]rune(u.Email))
	if emailLen < 3 || emailLen > 100 {
		return fmt.Errorf("invalid `Email` len: %d: %w", emailLen, core_errors.ErrInvalidArgument)
	}
	return nil
}

func NewUserUnitialized(
	email string,
	passwordHash string,
	role string,
) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		email,
		passwordHash,
		role,
	)
}

type UserPatch struct {
	Email    Nullable[string]
	Password Nullable[string]
	Role     Nullable[enum.Role]
}

func NewUserPatch(
	email Nullable[string],
	password Nullable[string],
	role Nullable[enum.Role],
) UserPatch {
	return UserPatch{
		Email:    email,
		Password: password,
		Role:     role,
	}
}

func (p *UserPatch) Validate() error {
	if p.Email.Set && p.Email.Value == nil {
		return fmt.Errorf("`Email` can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	if p.Password.Set && p.Password.Value == nil {
		return fmt.Errorf("'Password' can't be patched to NULL : %w", core_errors.ErrInvalidArgument)
	}

	if p.Role.Set && p.Role.Value == nil {
		return fmt.Errorf("'Role' can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}
