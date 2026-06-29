package groups_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

type GetGroupsResponse []GroupsDTOResponse

func (h *GroupsHTTPHandler) GetGroups(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := core_http_request.GetLimitOffsetQueryParams(c)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get 'limit'/'offset' query params",
		)
		return
	}

	groupsDomains, err := h.groupsService.GetGroups(ctx, limit, offset)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get groups",
		)
		return
	}
	response := GetGroupsResponse(groupsDTOFromDomains(groupsDomains))

	c.JSON(http.StatusOK, response)

}
