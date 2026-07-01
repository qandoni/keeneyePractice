package auth_jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	core_auth "github.com/qandoni/keeneyePractice/internal/core/auth"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

type JWTManager struct {
	secretKey string
}

func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{
		secretKey: secretKey,
	}
}

func (m *JWTManager) GenerateAccessToken(
	user domain.User,
) (string, error) {
	claims := Claims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(1 * time.Minute),
			),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(
		[]byte(m.secretKey),
	)
}

func (m *JWTManager) ParseAccessToken(
	tokenString string,
) (core_auth.AuthInfo, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwt.Token) (any, error) {
			return []byte(m.secretKey), nil
		},
	)
	if err != nil {
		return core_auth.AuthInfo{}, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return core_auth.AuthInfo{}, core_errors.ErrInvalidToken
	}
	return core_auth.AuthInfo{
		UserID: claims.UserID,
		Role:   claims.Role,
	}, nil
}
