package teachers_transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

type CreateTeacherRequest struct {
	UserID      int    `json:"user_id" validate:"required"`
	FIO         string `json:"fio" validate:"required,min=3,max=100"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=15,startswith=+"`
}

type CreateTeacherResponse TeacherDTOResponse

func (h *TeachersHTTPHandler) CreateTeacher(c *gin.Context) {
	ctx := c.Request.Context()

	var request CreateTeacherRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	teacher := domainFromDTO(request)

	teacher, err := h.teachersService.CreateTeacher(ctx, teacher)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to create teacher",
		)
		return
	}
	response := CreateTeacherResponse(
		teacherDTOFromDomain(teacher),
	)
	c.JSON(http.StatusCreated, response)
}

func domainFromDTO(dto CreateTeacherRequest) domain.Teacher {
	return domain.NewTeacherUninitialized(dto.UserID, dto.FIO,
		dto.PhoneNumber)
}
