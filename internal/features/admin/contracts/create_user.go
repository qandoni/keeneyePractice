package admin_contracts

import "github.com/qandoni/keeneyePractice/internal/core/enum"

type CreateUserCommand struct {
	Email       string
	Password    string
	Role        enum.Role
	FIO         string
	PhoneNumber string
	GroupID     *int
}
