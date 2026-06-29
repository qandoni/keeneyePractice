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
}
