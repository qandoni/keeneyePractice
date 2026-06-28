package auth_service

import (
	"context"

	core_auth "github.com/qandoni/keeneyePractice/internal/core/auth"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

type AuthService struct {
	usersRepository UsersRepository
	passwordHasher  PasswordHasher
	jwtManager      JWTManager
}

type UsersRepository interface {
	GetUserByLogin(
		ctx context.Context,
		login string,
	) (domain.User, error)
}

type PasswordHasher interface {
	Hash(
		password string,
	) (string, error)
	Compare(
		hash string,
		password string,
	) error
}

type JWTManager interface {
	GenerateToken(
		user domain.User,
	) (string, error)
	ParseToken(
		token string,
	) (core_auth.AuthInfo, error)
}

func NewAuthService(
	usersRepository UsersRepository,
	passwordHasher PasswordHasher,
	jwtManager JWTManager,
) *AuthService {
	return &AuthService{
		usersRepository: usersRepository,
		passwordHasher:  passwordHasher,
		jwtManager:      jwtManager,
	}
}
