package students_transport_http

import (
	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

func (h *StudentsHTTPHandler) DeleteStudent(c *gin.Context) {
	ctx := c.Request.Context()

	studentID, err := core_http_request.GetIntPathValue(c, "id")
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get int path value",
		)
		return
	}
	if err := h.studentsService.DeleteStudent(ctx, studentID); err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to delete student",
		)
		return
	}
	core_http_response.RespondNoContent(c)
}
