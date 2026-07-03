package students_transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
)

func (h *StudentsHTTPHandler) DeleteStudent(c *gin.Context) {
	ctx := c.Request.Context()

	studentID, err := core_http_request.GetIntPathValue(c, "id")
	if err != nil {
		c.Error(err).SetMeta("failed to get int path value")
		return
	}
	if err := h.studentsService.DeleteStudent(ctx, studentID); err != nil {
		c.Error(err).SetMeta("failed to delete student")
		return
	}
	c.Status(http.StatusNoContent)
}
