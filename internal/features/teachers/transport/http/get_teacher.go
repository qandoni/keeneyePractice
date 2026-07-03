package teachers_transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
)

type GetTeacherResponse TeacherDTOResponse

func (h *TeachersHTTPHandler) GetTeacher(c *gin.Context) {
	ctx := c.Request.Context()

	teacherID, err := core_http_request.GetIntPathValue(c, "id")

	if err != nil {
		c.Error(err).SetMeta("failed to get 'teacherID' path value")
		return
	}
	teacher, err := h.teachersService.GetTeacher(ctx, teacherID)
	if err != nil {
		c.Error(err).SetMeta("failed to get teacher")
		return
	}
	response := GetTeacherResponse(teacherDTOFromDomain(teacher))
	c.JSON(http.StatusOK, response)

}
