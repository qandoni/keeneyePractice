package students_transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

type CreateStudentRequest struct {
	UserID      int    `json:"user_id" validate:"required"`
	GroupID     int    `json:"group_id" validate:"required"`
	FIO         string `json:"fio" validate:"required,min=3,max=100"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=15,startswith=+"`
}

type CreateStudentResponse StudentDTOResponse

func (h *StudentsHTTPHandler) CreateStudent(c *gin.Context) {
	ctx := c.Request.Context()

	var request CreateStudentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err).SetMeta("failed to decode and validate HTTP request")
		return
	}

	student := domainFromDTO(request)

	student, err := h.studentsService.CreateStudent(
		ctx,
		student,
	)
	if err != nil {
		c.Error(err).SetMeta("failed to create student")
		return
	}

	response := CreateStudentResponse(
		studentDTOFromDomain(student),
	)

	c.JSON(
		http.StatusCreated,
		response,
	)
}

func domainFromDTO(dto CreateStudentRequest) domain.Student {
	return domain.NewStudentUninitialized(dto.UserID, dto.GroupID, dto.FIO, dto.PhoneNumber)
}
