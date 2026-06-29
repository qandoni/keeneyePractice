package teachers_transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

type GetTeacherResponse TeacherDTOResponse

func (h *TeachersHTTPHandler) GetTeacher(c *gin.Context) {
	ctx := c.Request.Context()

	teacherID, err := core_http_request.GetIntPathValue(c, "id")

	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get 'teacherID' path value",
		)
		return
	}
	teacher, err := h.teachersService.GetTeacher(ctx, teacherID)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get teacher",
		)
		return
	}
	response := GetTeacherResponse(teacherDTOFromDomain(teacher))
	c.JSON(http.StatusOK, response)

}
