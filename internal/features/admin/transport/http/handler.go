package admin_transport_http

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
	admin_contracts "github.com/qandoni/keeneyePractice/internal/features/admin/contracts"
)

type AdminHTTPHandler struct {
	adminService AdminService
}

type AdminService interface {
	CreateUser(
		ctx context.Context,
		cmd admin_contracts.CreateUserCommand,
	) (domain.User, error)
	AssignTeacherToGroup(
		ctx context.Context,
		teacherID int,
		groupID int,
	) error
	RemoveTeacherFromGroup(
		ctx context.Context,
		teacherID int,
		groupID int,
	) error
	CreateGroup(
		ctx context.Context,
		name string,
	) (domain.Group, error)
	GetTeachers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.Teacher, error)
}

func NewAdminHTTPHandler(
	adminService AdminService,
) *AdminHTTPHandler {
	return &AdminHTTPHandler{
		adminService: adminService,
	}
}
func (h *AdminHTTPHandler) Register(rg *gin.RouterGroup) {

	users := rg.Group("/users")
	users.POST("", h.CreateUser)

	teachers := rg.Group("/teachers")
	teachers.GET("", h.GetTeachers)
	teachers.POST("/:teacher_id/groups/:group_id", h.AssignTeacherToGroup)
	teachers.DELETE("/:teacher_id/groups/:group_id", h.RemoveTeacherFromGroup)

	groups := rg.Group("/groups")
	groups.POST("", h.CreateGroup)
}
