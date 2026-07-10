package registration_service

import (
	"context"
	"fmt"
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	registration_contracts "github.com/qandoni/keeneyePractice/internal/features/registration_requests/contracts"
)

func (s *RegistrationRequestService) Import(
	ctx context.Context,
	input registration_contracts.ImportInput,
) error {
	return s.txManager.WithinTransaction(ctx, func(ctx context.Context) error {

		for _, row := range input.Rows {

			if row.Role != enum.RoleStudent && row.Role != enum.RoleTeacher {
				return fmt.Errorf("invalid role: %s: %w", row.Role, core_errors.ErrInvalidArgument)
			}

			var groupID *int
			if row.Role == enum.RoleStudent {
				group, err := s.groupsService.GetGroupByName(ctx, row.Group)
				if err != nil {
					return fmt.Errorf("group with name '%s': %w", row.Group, core_errors.ErrNotFound)
				}
				groupID = &group.ID
			}

			token, err := s.tokenGenerator.Generate()
			if err != nil {
				return fmt.Errorf("generate token: %w", err)
			}

			tokenHash := s.sha256Hasher.Hash(token)

			req := domain.RegistrationRequest{
				FIO:              row.FIO,
				Email:            row.Email,
				PhoneNumber:      row.PhoneNumber,
				Role:             row.Role,
				GroupID:          groupID,
				TokenHash:        tokenHash,
				Status:           enum.StatusPending,
				EmailStatus:      enum.StatusEmailPending,
				EmailRetryCount:  0,
				LastEmailAttempt: nil,
				EmailSentAt:      nil,
				ExpiresAt:        time.Now().Add(24 * time.Hour),
			}

			if err := s.registrationRepository.Create(ctx, req); err != nil {
				return fmt.Errorf("create request: %w", err)
			}
		}

		return nil
	})
}
