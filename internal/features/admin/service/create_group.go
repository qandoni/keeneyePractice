package admin_service

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	group_contracts "github.com/qandoni/keeneyePractice/internal/features/groups/contracts"
)

func (s *AdminService) CreateGroup(
	ctx context.Context,
	name string,
) (domain.Group, error) {

	return s.groupsService.CreateGroup(ctx, group_contracts.CreateGroupInput{
		Name: name,
	})
}
