package registration_worker

import (
	"context"
	"fmt"
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	registration_contracts "github.com/qandoni/keeneyePractice/internal/features/registration_requests/contracts"
)

type RetryEmailWorker struct {
	repository           RegistrationRequestRepository
	emailDispatchService EmailDispatchService
}

type EmailDispatchService interface {
	SendRegistrationEmail(
		ctx context.Context,
		requestID int,
	) error
}

type RegistrationRequestRepository interface {
	GetRetryableEmailRequests(
		ctx context.Context,
	) ([]domain.RegistrationRequest, error)

	UpdateAfterEmailAttempt(
		ctx context.Context,
		input registration_contracts.UpdateEmailAttemptInput,
	) error
}

func NewRetryEmailWorker(
	repository RegistrationRequestRepository,
	emailDispatchService EmailDispatchService,
) *RetryEmailWorker {
	return &RetryEmailWorker{
		repository:           repository,
		emailDispatchService: emailDispatchService,
	}
}

func (w *RetryEmailWorker) Run(ctx context.Context) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			err := w.retry(ctx)
			if err != nil {
				return fmt.Errorf("method retry: %w", err)
			}
		}
	}
}

func (w *RetryEmailWorker) retry(ctx context.Context) error {
	requests, err := w.repository.GetRetryableEmailRequests(ctx)
	if err != nil {
		return err
	}
	for _, req := range requests {
		if err := w.emailDispatchService.SendRegistrationEmail(
			ctx,
			req.ID,
		); err != nil {
			return fmt.Errorf("send registration email: %w", err)
		}
	}

	return nil
}
