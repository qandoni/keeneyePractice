package users_contracts

import (
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

type CreateUserInput struct {
	Email    string
	Password string
	Role     enum.Role
}

func (i *CreateUserInput) Validate() error {
	emailLen := len([]rune(i.Email))
	if emailLen < 3 || emailLen > 100 {
		return fmt.Errorf(
			"invalid 'Email' len: %d: %w",
			emailLen,
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}
