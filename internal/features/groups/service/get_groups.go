package groups_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *GroupsService) GetGroups(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Group, error) {

	groups, err := s.groupsRepository.GetGroups(
		ctx,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"get groups: %w",
			err,
		)
	}

	return groups, nil
}
