package email_dispatch_service

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	registration_contracts "github.com/qandoni/keeneyePractice/internal/features/registration_requests/contracts"
)

type EmailDispatchService struct {
	registrationRepository RegistrationRequestRepository
	emailSender            EmailSender
	tokenGenerator         TokenGenerator
	sha256Hasher           Sha256Hasher
	config                 Config
}

type Producer interface {
	PublishRegistrationRequest(
		ctx context.Context,
		requestID int,
	) error
	PublishRetryRegistrationRequest(
		ctx context.Context,
		requestID int,
		retry int,
	) error
	PublishDeadRegistrationRequest(
		ctx context.Context,
		requestID int,
	) error
}

type RegistrationRequestRepository interface {
	GetByID(
		ctx context.Context,
		id int,
	) (domain.RegistrationRequest, error)

	UpdateAfterEmailAttempt(
		ctx context.Context,
		input registration_contracts.UpdateEmailAttemptInput,
	) error
}

type EmailSender interface {
	SendRegistrationEmail(
		ctx context.Context,
		to string,
		subject string,
		body string,
	) error
}

type TokenGenerator interface {
	Generate() (string, error)
}

type Sha256Hasher interface {
	Hash(string) string
}

func NewEmailDispatchService(
	registrationRepository RegistrationRequestRepository,
	emailSender EmailSender,
	tokenGenerator TokenGenerator,
	sha256Hasher Sha256Hasher,
	config Config,
) *EmailDispatchService {

	return &EmailDispatchService{
		registrationRepository: registrationRepository,
		emailSender:            emailSender,
		tokenGenerator:         tokenGenerator,
		sha256Hasher:           sha256Hasher,
		config:                 config,
	}
}
