package users_contracts

import (
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

type CreateUserInput struct {
	Login    string
	Password string
	Role     enum.Role
}

func (i *CreateUserInput) Validate() error {
	loginLen := len([]rune(i.Login))
	if loginLen < 3 || loginLen > 100 {
		return fmt.Errorf(
			"invalid 'Login' len: %d: %w",
			loginLen,
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}
