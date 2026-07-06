package groups_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *GroupsService) GetGroupByName(
	ctx context.Context,
	name string,
) (domain.Group, error) {
	group, err := s.groupsRepository.GetGroupByName(ctx, name)
	if err != nil {
		return domain.Group{}, fmt.Errorf("get group by name: %w", err)
	}
	return group, nil
}
