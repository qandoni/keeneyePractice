package admin_transport_http

import "github.com/qandoni/keeneyePractice/internal/core/domain"

type AdminDTOResponse struct {
	ID           int    `json:"ID"`
	Version      int    `json:"version"`
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}

func AdminDTOFromDomain(user domain.User) AdminDTOResponse {
	return AdminDTOResponse{
		ID:           user.ID,
		Version:      user.Version,
		Login:        user.Login,
		PasswordHash: user.PasswordHash,
		Role:         string(user.Role),
	}
}
