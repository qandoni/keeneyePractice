package auth_service

import (
	"context"

	core_auth "github.com/qandoni/keeneyePractice/internal/core/auth"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_postgres "github.com/qandoni/keeneyePractice/internal/core/repository/postgres"
)

type AuthService struct {
	usersRepository   UsersRepository
	refreshRepository RefreshTokenRepository
	passwordHasher    PasswordHasher
	sha256Hasher      Sha256Hasher
	jwtManager        JWTManager
	refreshGenerator  RefreshGenerator
	txManager         core_postgres.TransactionManager
}

type UsersRepository interface {
	GetUserByLogin(
		ctx context.Context,
		login string,
	) (domain.User, error)
	GetUser(
		ctx context.Context,
		id int,
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

type Sha256Hasher interface {
	Hash(value string) string
}

type JWTManager interface {
	GenerateAccessToken(
		user domain.User,
	) (string, error)
	ParseAccessToken(
		token string,
	) (core_auth.AuthInfo, error)
}

type RefreshGenerator interface {
	Generate() (string, error)
	Hash(
		token string,
	) string
}

type RefreshTokenRepository interface {
	Create(
		ctx context.Context,
		token domain.RefreshToken,
	) error
	Get(
		ctx context.Context,
		tokenHash string,
	) (domain.RefreshToken, error)
	Delete(
		ctx context.Context,
		tokenHash string,
	) error
	PatchRefreshToken(
		ctx context.Context,
		token domain.RefreshToken,
	) (domain.RefreshToken, error)
}

func NewAuthService(
	usersRepository UsersRepository,
	refreshRepository RefreshTokenRepository,
	passwordHasher PasswordHasher,
	sha256Hasher Sha256Hasher,
	jwtManager JWTManager,
	refreshGenerator RefreshGenerator,
	txManager core_postgres.TransactionManager,
) *AuthService {
	return &AuthService{
		usersRepository:   usersRepository,
		refreshRepository: refreshRepository,
		passwordHasher:    passwordHasher,
		sha256Hasher:      sha256Hasher,
		jwtManager:        jwtManager,
		refreshGenerator:  refreshGenerator,
		txManager:         txManager,
	}
}
