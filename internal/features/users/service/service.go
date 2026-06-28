package users_service

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
	passwordHasher  PasswordHasher
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	GetUserByLogin(
		ctx context.Context,
		login string,
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
		user domain.User,
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

func NewUsersService(
	usersRepository UsersRepository,
	passwordHasher PasswordHasher,
) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
		passwordHasher:  passwordHasher,
	}
}
