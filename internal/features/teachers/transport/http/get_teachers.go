package teachers_transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
)

type GetTeachersResponse []TeacherDTOResponse

func (h *TeachersHTTPHandler) GetTeachers(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := core_http_request.GetLimitOffsetQueryParams(c)
	if err != nil {
		c.Error(err).SetMeta("failed to get 'limit'/'offset' query params")
		return
	}

	teacherDomains, err := h.teachersService.GetTeachers(ctx, limit, offset)
	if err != nil {
		c.Error(err).SetMeta("failed to get teachers")
		return
	}
	response := GetTeachersResponse(teachersDTOFromDomains(teacherDomains))

	c.JSON(http.StatusOK, response)

}
