package users_http_transport

import "github.com/qandoni/keeneyePractice/internal/core/domain"

type UserDTOResponse struct {
	ID           int
	Version      int
	Login        string
	PasswordHash string
	Role         string
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:           user.ID,
		Version:      user.Version,
		Login:        user.Login,
		PasswordHash: user.PasswordHash,
		Role:         string(user.Role),
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))
	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}
	return usersDTO
}
