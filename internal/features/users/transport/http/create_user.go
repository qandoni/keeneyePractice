package users_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
	users_contracts "github.com/qandoni/keeneyePractice/internal/features/users/contracts"
)

type CreateUserRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type CreateUserResponse UserDTOResponse

func (h *UsersHTTPHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	input := users_contracts.CreateUserInput{
		Login:    request.Login,
		Password: request.Password,
		Role:     enum.Role(request.Role),
	}

	user, err := h.usersService.CreateUser(ctx, input)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to create user",
		)
		return
	}

	response := CreateUserResponse(userDTOFromDomain(user))
	c.JSON(http.StatusCreated, response)
}
