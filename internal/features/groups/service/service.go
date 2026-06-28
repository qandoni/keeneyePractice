package groups_service

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

type GroupsService struct {
	groupsRepository GroupsRepository
}

type GroupsRepository interface {
	CreateGroup(ctx context.Context, group domain.Group) (domain.Group, error)

	GetGroup(ctx context.Context, id int) (domain.Group, error)

	GetGroups(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.Group, error)

	PatchGroup(
		ctx context.Context,
		group domain.Group,
	) (domain.Group, error)

	DeleteGroup(
		ctx context.Context,
		id int,
	) error
}

func NewGroupsService(
	groupsRepository GroupsRepository,
) *GroupsService {
	return &GroupsService{
		groupsRepository: groupsRepository,
	}
}
