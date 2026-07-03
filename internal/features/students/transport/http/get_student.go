package students_transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
)

type GetStudentResponse StudentDTOResponse

func (h *StudentsHTTPHandler) GetStudent(c *gin.Context) {
	ctx := c.Request.Context()

	studentID, err := core_http_request.GetIntPathValue(c, "id")
	if err != nil {
		c.Error(err).SetMeta("failed to get 'studentID' path value")
		return
	}
	student, err := h.studentsService.GetStudent(ctx, studentID)
	if err != nil {
		c.Error(err).SetMeta("failed to get student")
		return
	}
	response := GetStudentResponse(studentDTOFromDomain(student))

	c.JSON(http.StatusOK, response)

}
