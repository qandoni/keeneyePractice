package students_transport_http

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_http_middleware "github.com/qandoni/keeneyePractice/internal/core/transport/http/middleware"
)

type StudentsHTTPHandler struct {
	studentsService StudentsService
}

type StudentsService interface {
	CreateStudent(
		ctx context.Context,
		student domain.Student,
	) (domain.Student, error)
	GetStudent(
		ctx context.Context,
		id int,
	) (domain.Student, error)
	PatchStudent(
		ctx context.Context,
		id int,
		patch domain.StudentPatch,
	) (domain.Student, error)
	DeleteStudent(
		ctx context.Context,
		id int,
	) error
	GetStudents(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.Student, error)
}

func NewStudentsHTTPHandler(
	studentsService StudentsService,
) *StudentsHTTPHandler {
	return &StudentsHTTPHandler{
		studentsService: studentsService,
	}
}

func (h *StudentsHTTPHandler) Register(rg *gin.RouterGroup) {
	students := rg.Group("")

	students.GET("", h.GetStudents)
	students.GET("/:id", h.GetStudent)

	students.PATCH("/:id", h.PatchStudent)
	students.DELETE("/:id", core_http_middleware.Role(enum.RoleAdmin), h.DeleteStudent)
}
