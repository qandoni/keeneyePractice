package teachers_transport_http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *TeachersHTTPHandler) AssignToGroup(c *gin.Context) {
	ctx := c.Request.Context()

	teacherID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err).SetMeta("invalid teacher id")
		return
	}

	groupID, err := strconv.Atoi(c.Param("group_id"))
	if err != nil {
		c.Error(err).SetMeta("invalid group id")
		return
	}

	err = h.teachersService.AssignToGroup(ctx, teacherID, groupID)
	if err != nil {
		c.Error(err).SetMeta("failed to assign teacher to group")
		return
	}

	c.Status(http.StatusNoContent)
}
