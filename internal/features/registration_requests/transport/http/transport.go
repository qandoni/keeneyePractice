package registration_http_transport

import (
	"context"

	registration_contracts "github.com/qandoni/keeneyePractice/internal/features/registration_requests/contracts"
)

type RegistrationRequestsHTTPHandler struct {
	service RegistrationRequestsService
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
}

func NewRegistrationRequestsHTTPHandler(
	service RegistrationRequestsService,
) *RegistrationRequestsHTTPHandler {
	return &RegistrationRequestsHTTPHandler{
		service: service,
	}
}
