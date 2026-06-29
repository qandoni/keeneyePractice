package groups_http_transport

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

type GroupsHTTPHandler struct {
	groupsService GroupsService
}

type GroupsService interface {
	CreateGroup(
		ctx context.Context,
		group domain.Group,
	) (domain.Group, error)
	PatchGroup(
		ctx context.Context,
		id int,
		patch domain.GroupPatch,
	) (domain.Group, error)
	GetGroup(
		ctx context.Context,
		id int,
	) (domain.Group, error)
	GetGroups(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.Group, error)
	DeleteGroup(
		ctx context.Context,
		id int,
	) error
}

func NewGroupsHTTPHandler(
	groupsService GroupsService,
) *GroupsHTTPHandler {
	return &GroupsHTTPHandler{
		groupsService: groupsService,
	}
}

func (h *GroupsHTTPHandler) Register(rg *gin.RouterGroup) {
	groups := rg.Group("")
	groups.POST("", h.CreateGroup)
	groups.GET(":id", h.GetGroup)
	groups.GET("", h.GetGroups)
	groups.DELETE("/:id", h.DeleteGroup)
	groups.PATCH("/:id", h.PatchGroup)
}
