package groups_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

type CreateGroupRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateGroupResponse GroupsDTOResponse

func (h *GroupsHTTPHandler) CreateGroup(c *gin.Context) {
	ctx := c.Request.Context()

	var req CreateGroupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		core_http_response.RespondError(c, err, "invalid request")
		return
	}

	group := domainFromDTO(req)

	group, err := h.groupsService.CreateGroup(ctx, group)
	if err != nil {
		core_http_response.RespondError(c, err, "failed to create group")
		return
	}
	response := CreateGroupResponse(groupDTOFromDomain(group))
	c.JSON(http.StatusCreated, response)
}

func domainFromDTO(dto CreateGroupRequest) domain.Group {
	return domain.NewGroupUninitialized(dto.Name)
}
