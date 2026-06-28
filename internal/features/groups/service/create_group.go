package groups_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	group_contracts "github.com/qandoni/keeneyePractice/internal/features/groups/contracts"
)

func (s *GroupsService) CreateGroup(
	ctx context.Context,
	input group_contracts.CreateGroupInput,
) (domain.Group, error) {

	if err := input.Validate(); err != nil {
		return domain.Group{}, fmt.Errorf(
			"validate create group input: %w",
			err,
		)
	}

	group := domain.NewGroupUninitialized(
		input.Name,
	)

	group, err := s.groupsRepository.CreateGroup(
		ctx,
		group,
	)
	if err != nil {
		return domain.Group{}, fmt.Errorf(
			"create group: %w",
			err,
		)
	}

	return group, nil
}
