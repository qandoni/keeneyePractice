package teachers_transport_http

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

type TeachersHTTPHandler struct {
	teachersService TeachersService
}

type TeachersService interface {
	CreateTeacher(
		ctx context.Context,
		teacher domain.Teacher,
	) (domain.Teacher, error)
	PatchTeacher(
		ctx context.Context,
		id int,
		patch domain.TeacherPatch,
	) (domain.Teacher, error)
	GetTeacher(
		ctx context.Context,
		id int,
	) (domain.Teacher, error)
	GetTeachers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.Teacher, error)
	DeleteTeacher(
		ctx context.Context,
		id int,
	) error
	AssignToGroup(
		ctx context.Context,
		teacherID int,
		groupID int,
	) error
	RemoveFromGroup(
		ctx context.Context,
		teacherID int,
		groupID int,
	) error
}

func NewTeachersHTTPHandler(
	teachersService TeachersService,
) *TeachersHTTPHandler {
	return &TeachersHTTPHandler{
		teachersService: teachersService,
	}
}

func (h *TeachersHTTPHandler) Register(rg *gin.RouterGroup) {
	teachers := rg.Group("")
	teachers.POST("", h.CreateTeacher)
	teachers.GET("", h.GetTeachers)
	teacher := rg.Group("/:id")
	{
		teacher.GET("", h.GetTeacher)
		teacher.PATCH("", h.PatchTeacher)
		teacher.DELETE("", h.DeleteTeacher)

		teacher.POST("/groups/:group_id", h.AssignToGroup)
		teacher.DELETE("/groups/:group_id", h.RemoveFromGroup)
	}
}
