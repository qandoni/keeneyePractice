package domain

import (
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/enum"
)

type RegistrationRequest struct {
	ID               int
	Version          int
	FIO              string
	Email            string
	PhoneNumber      string
	Role             enum.Role
	GroupID          *int
	TokenHash        string
	Status           enum.RegistrationStatus
	EmailStatus      enum.EmailStatus
	EmailRetryCount  int
	LastEmailAttempt *time.Time
	EmailSentAt      *time.Time
	ExpiresAt        time.Time
	CreatedAt        time.Time
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
	status enum.RegistrationStatus,
	emailStatus enum.EmailStatus,
	emailRetryCount int,
	lastEmailAttempt *time.Time,
	emailSentAt *time.Time,
	expiresAt time.Time,
	createdAt time.Time,
) RegistrationRequest {
	return RegistrationRequest{
		ID:               id,
		Version:          version,
		FIO:              fio,
		Email:            email,
		PhoneNumber:      phoneNumber,
		Role:             role,
		GroupID:          groupID,
		TokenHash:        tokenHash,
		Status:           status,
		EmailStatus:      emailStatus,
		EmailRetryCount:  emailRetryCount,
		LastEmailAttempt: lastEmailAttempt,
		EmailSentAt:      emailSentAt,
		ExpiresAt:        expiresAt,
		CreatedAt:        createdAt,
	}
}
