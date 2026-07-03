package teachers_transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
)

func (h *TeachersHTTPHandler) DeleteTeacher(c *gin.Context) {
	ctx := c.Request.Context()

	teacherID, err := core_http_request.GetIntPathValue(c, "id")
	if err != nil {
		c.Error(err).SetMeta("failed to get int path value")
		return
	}
	if err := h.teachersService.DeleteTeacher(ctx, teacherID); err != nil {
		c.Error(err).SetMeta("failed to delete teacher")
		return
	}
	c.Status(http.StatusNoContent)
}
