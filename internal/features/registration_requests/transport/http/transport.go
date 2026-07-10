package registration_http_transport

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	registration_contracts "github.com/qandoni/keeneyePractice/internal/features/registration_requests/contracts"
)

type RegistrationRequestsHTTPHandler struct {
	registrationService RegistrationRequestsService
}

type RegistrationRequestsService interface {
	Import(
		ctx context.Context,
		input registration_contracts.ImportInput,
	) error
	Complete(
		ctx context.Context,
		input registration_contracts.CompleteInput,
	) error
	GetAll(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.RegistrationRequest, error)
}

func NewRegistrationRequestsHTTPHandler(
	service RegistrationRequestsService,
) *RegistrationRequestsHTTPHandler {
	return &RegistrationRequestsHTTPHandler{
		registrationService: service,
	}
}
