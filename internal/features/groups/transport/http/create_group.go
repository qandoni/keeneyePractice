package groups_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

type CreateGroupRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateGroupResponse GroupsDTOResponse

func (h *GroupsHTTPHandler) CreateGroup(c *gin.Context) {
	ctx := c.Request.Context()

	var req CreateGroupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err).SetMeta("failed to decode and validate HTTP request")
		return
	}

	group := domainFromDTO(req)

	group, err := h.groupsService.CreateGroup(ctx, group)
	if err != nil {
		c.Error(err).SetMeta("failed to create group")
		return
	}
	response := CreateGroupResponse(groupDTOFromDomain(group))
	c.JSON(http.StatusCreated, response)
}

func domainFromDTO(dto CreateGroupRequest) domain.Group {
	return domain.NewGroupUninitialized(dto.Name)
}
