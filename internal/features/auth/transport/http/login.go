package auth_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
	auth_contracts "github.com/qandoni/keeneyePractice/internal/features/auth/contracts"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *AuthHTTPHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	input := auth_contracts.LoginInput{
		Login:    request.Login,
		Password: request.Password,
	}
	token, err := h.authService.Login(
		ctx,
		input,
	)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to login",
		)
		return
	}
	response := LoginResponse{
		Token: token,
	}

	c.JSON(http.StatusOK, response)
}
