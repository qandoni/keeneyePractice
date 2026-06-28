package admin_transport_http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

func (h *AdminHTTPHandler) RemoveTeacherFromGroup(c *gin.Context) {
	ctx := c.Request.Context()

	teacherID, err := strconv.Atoi(c.Param("teacher_id"))
	if err != nil {
		core_http_response.RespondError(c, err, "invalid teacher id")
		return
	}

	groupID, err := strconv.Atoi(c.Param("group_id"))
	if err != nil {
		core_http_response.RespondError(c, err, "invalid group id")
		return
	}

	err = h.adminService.RemoveTeacherFromGroup(ctx, teacherID, groupID)
	if err != nil {
		core_http_response.RespondError(c, err, "failed to remove teacher from group")
		return
	}

	c.Status(http.StatusNoContent)
}
