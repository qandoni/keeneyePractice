package registration_postgres_repository

import (
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
)

type RegistrationRequestModel struct {
	ID          int
	Version     int
	FIO         string
	Email       string
	PhoneNumber string
	Role        string
	GroupID     *int
	TokenHash   string
	ExpiresAt   time.Time
	Status      string
}

func modelToDomain(model RegistrationRequestModel) domain.RegistrationRequest {
	return domain.RegistrationRequest{
		ID:          model.ID,
		Version:     model.Version,
		FIO:         model.FIO,
		Email:       model.Email,
		PhoneNumber: model.PhoneNumber,
		Role:        enum.Role(model.Role),
		GroupID:     model.GroupID,
		TokenHash:   model.TokenHash,
		ExpiresAt:   model.ExpiresAt,
		Status:      enum.RegistrationStatus(model.Status),
	}
}
