package registration_contracts

import (
	"github.com/qandoni/keeneyePractice/internal/core/enum"
)

type ImportInput struct {
	Rows []ImportRow
}

type ImportRow struct {
	FIO         string
	Email       string
	PhoneNumber string
	Role        enum.Role
	Group       string
}
