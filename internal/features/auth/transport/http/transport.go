package auth_http_transport

import (
	"context"

	"github.com/gin-gonic/gin"
	auth_contracts "github.com/qandoni/keeneyePractice/internal/features/auth/contracts"
)

type AuthHTTPHandler struct {
	authService AuthService
}

type AuthService interface {
	Login(
		ctx context.Context,
		input auth_contracts.LoginInput,
	) (auth_contracts.LoginOutput, error)
	Refresh(
		ctx context.Context,
		input auth_contracts.RefreshInput,
	) (auth_contracts.RefreshOutput, error)
}

func NewAuthHTTPHandler(
	authService AuthService,
) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		authService: authService,
	}
}

func (h *AuthHTTPHandler) Register(rg *gin.RouterGroup) {
	rg.POST("/login", h.Login)
	rg.POST("/refresh", h.Refresh)
}
