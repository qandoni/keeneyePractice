package users_http_transport

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
	users_contracts "github.com/qandoni/keeneyePractice/internal/features/users/contracts"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		input users_contracts.CreateUserInput,
	) (domain.User, error)
	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	DeleteUser(
		ctx context.Context,
		id int,
	) error
	PatchUser(
		ctx context.Context,
		id int,
		patch domain.UserPatch,
	) (domain.User, error)
}

func NewUsersHTTPHandler(
	usersService UsersService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

func (h *UsersHTTPHandler) Register(rg *gin.RouterGroup) {
	users := rg.Group("")

	users.GET("/:id", h.GetUser)
	users.GET("", h.GetUsers)
	users.PATCH("/:id", h.PatchUser)
	users.DELETE("/:id", h.DeleteUser)
}
