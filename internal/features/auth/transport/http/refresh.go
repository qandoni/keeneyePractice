package auth_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
	auth_contracts "github.com/qandoni/keeneyePractice/internal/features/auth/contracts"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *AuthHTTPHandler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()

	var request RefreshRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}
	input := auth_contracts.RefreshInput{
		RefreshToken: request.RefreshToken,
	}
	output, err := h.authService.Refresh(ctx, input)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to refresh token",
		)
	}
	response := RefreshResponse{
		AccessToken:  output.AccessToken,
		RefreshToken: output.RefreshToken,
	}
	c.JSON(http.StatusOK, response)
}
