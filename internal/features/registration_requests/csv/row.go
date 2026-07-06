package registration_csv

import "github.com/qandoni/keeneyePractice/internal/core/enum"

type Row struct {
	FIO         string
	Email       string
	PhoneNumber string
	Role        enum.Role
	Group       string
}
