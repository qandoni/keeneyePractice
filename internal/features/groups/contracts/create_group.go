package group_contracts

import (
	"fmt"

	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

type CreateGroupInput struct {
	Name string
}

func (i *CreateGroupInput) Validate() error {
	nameLen := len([]rune(i.Name))

	if nameLen < 2 || nameLen > 100 {
		return fmt.Errorf(
			"invalid 'Name' len: %d: %w",
			nameLen,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
