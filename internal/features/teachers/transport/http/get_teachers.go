package teachers_transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

type GetTeachersResponse []TeacherDTOResponse

func (h *TeachersHTTPHandler) GetTeachers(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := core_http_request.GetLimitOffsetQueryParams(c)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get 'limit'/'offset' query params",
		)
		return
	}

	teacherDomains, err := h.teachersService.GetTeachers(ctx, limit, offset)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get teachers",
		)
		return
	}
	response := GetTeachersResponse(teachersDTOFromDomains(teacherDomains))

	c.JSON(http.StatusOK, response)

}
