package groups_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
)

type GetGroupsResponse []GroupsDTOResponse

func (h *GroupsHTTPHandler) GetGroups(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := core_http_request.GetLimitOffsetQueryParams(c)
	if err != nil {
		c.Error(err).SetMeta("failed to get 'limit'/'offset' query params")
		return
	}

	groupsDomains, err := h.groupsService.GetGroups(ctx, limit, offset)
	if err != nil {
		c.Error(err).SetMeta("failed to get groups")
		return
	}
	response := GetGroupsResponse(groupsDTOFromDomains(groupsDomains))

	c.JSON(http.StatusOK, response)

}
