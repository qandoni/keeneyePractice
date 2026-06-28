package auth_jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
)

type Claims struct {
	UserID int       `json:"user_id"`
	Role   enum.Role `json:"role"`
	jwt.RegisteredClaims
}
