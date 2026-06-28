package domain

import (
	"fmt"

	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

type Group struct {
	ID      int
	Version int
	Name    string
}

func NewGroup(
	id int,
	version int,
	name string,
) Group {
	return Group{
		ID:      id,
		Version: version,
		Name:    name,
	}
}
func NewGroupUninitialized(
	name string,
) Group {
	return NewGroup(
		UninitializedGroupID,
		UninitializedVersion,
		name,
	)
}
func (g *Group) Validate() error {
	nameLen := len([]rune(g.Name))

	if nameLen < 2 || nameLen > 100 {
		return fmt.Errorf(
			"invalid 'Name' len: %d: %w",
			nameLen,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

type GroupPatch struct {
	Name Nullable[string]
}

func NewGroupPatch(
	name Nullable[string],
) GroupPatch {
	return GroupPatch{
		Name: name,
	}
}

func (p *GroupPatch) Validate() error {
	if p.Name.Set {
		if p.Name.Value == nil {
			return fmt.Errorf(
				"'Name' can't be NULL: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		nameLen := len([]rune(*p.Name.Value))
		if nameLen < 2 || nameLen > 100 {
			return fmt.Errorf(
				"invalid 'Name' len: %d: %w",
				nameLen,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

func (g *Group) ApplyPatch(
	patch GroupPatch,
) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate group patch: %w", err)
	}

	tmp := *g

	if patch.Name.Set {
		tmp.Name = *patch.Name.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched group: %w", err)
	}

	*g = tmp

	return nil
}
