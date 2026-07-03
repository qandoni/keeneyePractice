package groups_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
)

type GetGroupResponse GroupsDTOResponse

func (h *GroupsHTTPHandler) GetGroup(c *gin.Context) {
	ctx := c.Request.Context()

	groupID, err := core_http_request.GetIntPathValue(c, "id")

	if err != nil {
		c.Error(err).SetMeta("failed to get 'groupID' path value")
		return
	}
	group, err := h.groupsService.GetGroup(ctx, groupID)
	if err != nil {
		c.Error(err).SetMeta("failed to get group")
		return
	}
	response := GetGroupResponse(groupDTOFromDomain(group))
	c.JSON(http.StatusOK, response)

}
