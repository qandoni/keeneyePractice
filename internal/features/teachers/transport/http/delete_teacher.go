package teachers_transport_http

import (
	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

func (h *TeachersHTTPHandler) DeleteTeacher(c *gin.Context) {
	ctx := c.Request.Context()

	teacherID, err := core_http_request.GetIntPathValue(c, "id")
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get int path value",
		)
		return
	}
	if err := h.teachersService.DeleteTeacher(ctx, teacherID); err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to delete teacher",
		)
		return
	}
	core_http_response.RespondNoContent(c)
}
