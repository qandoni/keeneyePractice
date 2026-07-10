package registration_worker

import (
	"context"
	"fmt"
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
	registration_contracts "github.com/qandoni/keeneyePractice/internal/features/registration_requests/contracts"
)

type RetryEmailWorker struct {
	repository     RegistrationRequestRepository
	emailSender    EmailSender
	tokenGenerator TokenGenerator
	sha256Hasher   Sha256Hasher
	config         Config
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

func NewRetryEmailWorker(
	repository RegistrationRequestRepository,
	emailSender EmailSender,
	tokenGenerator TokenGenerator,
	sha256Hasher Sha256Hasher,
	config Config,
) *RetryEmailWorker {

	return &RetryEmailWorker{
		repository:     repository,
		emailSender:    emailSender,
		tokenGenerator: tokenGenerator,
		sha256Hasher:   sha256Hasher,
		config:         config,
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

		if req.EmailRetryCount >= 10 {
			now := time.Now()
			err := w.repository.UpdateAfterEmailAttempt(
				ctx,
				registration_contracts.UpdateEmailAttemptInput{
					ID:               req.ID,
					Version:          req.Version,
					TokenHash:        req.TokenHash,
					ExpiresAt:        req.ExpiresAt,
					EmailStatus:      enum.StatusEmailGiveUp,
					EmailRetryCount:  req.EmailRetryCount,
					LastEmailAttempt: &now,
					EmailSentAt:      req.EmailSentAt,
				},
			)

			if err != nil {
				return fmt.Errorf(
					"update give up status: %w",
					err,
				)
			}

			continue
		}

		token, err := w.tokenGenerator.Generate()
		if err != nil {
			return fmt.Errorf(
				"generate token: %w",
				err,
			)
		}

		tokenHash := w.sha256Hasher.Hash(token)

		expiresAt := time.Now().Add(24 * time.Hour)

		now := time.Now()

		err = w.emailSender.SendRegistrationEmail(
			ctx,
			req.Email,
			w.config.EmailSubject,
			fmt.Sprintf(
				"%s?token=%s",
				w.config.CompleteURL,
				token,
			),
		)

		input := registration_contracts.UpdateEmailAttemptInput{
			ID:               req.ID,
			Version:          req.Version,
			TokenHash:        tokenHash,
			ExpiresAt:        expiresAt,
			EmailStatus:      enum.StatusEmailSent,
			EmailRetryCount:  req.EmailRetryCount + 1,
			LastEmailAttempt: &now,
			EmailSentAt:      req.EmailSentAt,
		}

		if err != nil {
			input.EmailStatus = enum.StatusEmailFailed
			if input.EmailRetryCount >= 10 {
				input.EmailStatus = enum.StatusEmailGiveUp
			}
		} else {
			input.EmailStatus = enum.StatusEmailSent
			input.EmailSentAt = &now
		}

		if err := w.repository.UpdateAfterEmailAttempt(ctx, input); err != nil {
			return fmt.Errorf(
				"update after email attempt: %w",
				err,
			)
		}
	}

	return nil
}
