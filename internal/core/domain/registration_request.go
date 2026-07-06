package domain

import (
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/enum"
)

type RegistrationRequest struct {
	ID          int
	Version     int
	FIO         string
	Email       string
	PhoneNumber string
	Role        enum.Role
	GroupID     *int
	TokenHash   string
	ExpiresAt   time.Time
	Status      enum.RegistrationStatus
}

func NewRegistrationRequest(
	id int,
	version int,
	fio string,
	email string,
	phoneNumber string,
	role enum.Role,
	groupID *int,
	tokenHash string,
	expiresAt time.Time,
	status enum.RegistrationStatus,
) RegistrationRequest {
	return RegistrationRequest{
		ID:          id,
		Version:     version,
		FIO:         fio,
		Email:       email,
		PhoneNumber: phoneNumber,
		Role:        role,
		GroupID:     groupID,
		TokenHash:   tokenHash,
		ExpiresAt:   expiresAt,
		Status:      status,
	}
}
