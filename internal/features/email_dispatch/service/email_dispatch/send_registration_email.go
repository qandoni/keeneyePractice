package email_dispatch_service

import (
	"context"
	"fmt"
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/enum"
	registration_contracts "github.com/qandoni/keeneyePractice/internal/features/registration_requests/contracts"
)

func (s *EmailDispatchService) SendRegistrationEmail(
	ctx context.Context,
	requestID int,
) error {
	fmt.Println("SendRegistrationEmail called", requestID)
	req, err := s.registrationRepository.GetByID(ctx, requestID)
	if err != nil {
		return fmt.Errorf("get registration request: %w", err)
	}

	token, err := s.tokenGenerator.Generate()
	if err != nil {
		return fmt.Errorf("generate token: %w", err)

	}

	tokenHash := s.sha256Hasher.Hash(token)
	now := time.Now()
	expiresAt := now.Add(24 * time.Hour)

	err = s.emailSender.SendRegistrationEmail(
		ctx,
		req.Email,
		s.config.EmailSubject,
		fmt.Sprintf("%s?token=%s", s.config.CompleteURL, token),
	)

	input := registration_contracts.UpdateEmailAttemptInput{
		ID:               req.ID,
		Version:          req.Version,
		TokenHash:        tokenHash,
		ExpiresAt:        expiresAt,
		EmailRetryCount:  req.EmailRetryCount + 1,
		LastEmailAttempt: &now,
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

	if err := s.registrationRepository.UpdateAfterEmailAttempt(ctx, input); err != nil {
		return fmt.Errorf("update registration request: %w", err)

	}
	if err != nil {
		return fmt.Errorf("send email: %w", err)

	}

	return nil

}
