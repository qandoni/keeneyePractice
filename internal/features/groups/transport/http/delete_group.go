package groups_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
)

func (h *GroupsHTTPHandler) DeleteGroup(c *gin.Context) {
	ctx := c.Request.Context()

	groupID, err := core_http_request.GetIntPathValue(c, "id")
	if err != nil {
		c.Error(err).SetMeta("failed to get int path value")
		return
	}
	if err := h.groupsService.DeleteGroup(ctx, groupID); err != nil {
		c.Error(err).SetMeta("failed to delete group")
		return
	}
	c.Status(http.StatusNoContent)
}
