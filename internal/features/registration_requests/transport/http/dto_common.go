package registration_http_transport

import (
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
)

type RegistrationRequestDTOResponse struct {
	ID               int                     `json:"id"`
	Version          int                     `json:"version"`
	FIO              string                  `json:"fio"`
	Email            string                  `json:"email"`
	PhoneNumber      string                  `json:"phone_number"`
	Role             enum.Role               `json:"role"`
	GroupID          *int                    `json:"group_id"`
	TokenHash        string                  `json:"token_hash"`
	Status           enum.RegistrationStatus `json:"status"`
	EmailStatus      enum.EmailStatus        `json:"email_status"`
	EmailRetryCount  int                     `json:"email_retry_count"`
	LastEmailAttempt *time.Time              `json:"last_email_attempt"`
	EmailSentAt      *time.Time              `json:"email_sent_at"`
	ExpiresAt        time.Time               `json:"expires_at"`
	CreatedAt        time.Time               `json:"created_at"`
}

func NewRegistrationRequestDTOResponse(
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
) RegistrationRequestDTOResponse {
	return RegistrationRequestDTOResponse{
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

func RegistrationRequestDTOFromDomain(request domain.RegistrationRequest) RegistrationRequestDTOResponse {
	return RegistrationRequestDTOResponse{
		ID:               request.ID,
		Version:          request.Version,
		FIO:              request.FIO,
		Email:            request.Email,
		PhoneNumber:      request.PhoneNumber,
		Role:             request.Role,
		GroupID:          request.GroupID,
		TokenHash:        request.TokenHash,
		Status:           request.Status,
		EmailStatus:      request.EmailStatus,
		EmailRetryCount:  request.EmailRetryCount,
		LastEmailAttempt: request.LastEmailAttempt,
		EmailSentAt:      request.EmailSentAt,
		ExpiresAt:        request.ExpiresAt,
		CreatedAt:        request.CreatedAt,
	}
}

func RegistrationRequestDTOFromDomains(requests []domain.RegistrationRequest) []RegistrationRequestDTOResponse {
	requestsDTO := make([]RegistrationRequestDTOResponse, len(requests))
	for i, request := range requests {
		requestsDTO[i] = RegistrationRequestDTOFromDomain(request)
	}
	return requestsDTO
}
