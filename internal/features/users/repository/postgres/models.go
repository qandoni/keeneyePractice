package users_postgres_repository

import "github.com/qandoni/keeneyePractice/internal/core/domain"

type UserModel struct {
	ID           int
	Version      int
	Login        string
	PasswordHash string
	Role         string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))
	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.Login,
			user.PasswordHash,
			user.Role,
		)
	}
	return userDomains
}
