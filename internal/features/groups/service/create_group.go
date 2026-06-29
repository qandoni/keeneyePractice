package groups_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *GroupsService) CreateGroup(
	ctx context.Context,
	group domain.Group,
) (domain.Group, error) {
	if err := group.Validate(); err != nil {
		return domain.Group{}, fmt.Errorf("validate group domain: %w", err)
	}
	group, err := s.groupsRepository.CreateGroup(ctx, group)
	if err != nil {
		return domain.Group{}, fmt.Errorf("get group from repository: %w", err)
	}
	return group, nil
}
