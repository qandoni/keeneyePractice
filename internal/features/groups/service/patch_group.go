package groups_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *GroupsService) PatchGroup(
	ctx context.Context,
	id int,
	patch domain.GroupPatch,
) (domain.Group, error) {

	group, err := s.groupsRepository.GetGroup(
		ctx,
		id,
	)
	if err != nil {
		return domain.Group{}, fmt.Errorf(
			"get group: %w",
			err,
		)
	}

	if err := group.ApplyPatch(patch); err != nil {
		return domain.Group{}, fmt.Errorf(
			"apply patch: %w",
			err,
		)
	}

	group, err = s.groupsRepository.PatchGroup(
		ctx,
		group,
	)
	if err != nil {
		return domain.Group{}, fmt.Errorf(
			"patch group: %w",
			err,
		)
	}

	return group, nil
}
