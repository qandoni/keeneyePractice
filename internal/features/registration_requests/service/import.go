package registration_service

import (
	"context"
	"fmt"
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
	registration_contracts "github.com/qandoni/keeneyePractice/internal/features/registration_requests/contracts"
)

func (s *RegistrationRequestService) Import(
	ctx context.Context,
	input registration_contracts.ImportInput,
) error {

	type emailTask struct {
		email string
		token string
	}

	var tasks []emailTask

	err := s.txManager.WithinTransaction(ctx, func(ctx context.Context) error {

		for _, row := range input.Rows {

			if row.Role != enum.RoleStudent && row.Role != enum.RoleTeacher {
				return fmt.Errorf("invalid role: %s", row.Role)
			}

			_, err := s.usersService.GetUserByEmail(ctx, row.Email)
			if err == nil {
				return fmt.Errorf("user already exists: %s", row.Email)
			}

			var groupID *int
			if row.Role == enum.RoleStudent {
				group, err := s.groupsService.GetGroupByName(ctx, row.Group)
				if err != nil {
					return fmt.Errorf("group not found: %w", err)
				}
				groupID = &group.ID
			}

			token, err := s.tokenGenerator.Generate()
			if err != nil {
				return fmt.Errorf("generate token: %w", err)
			}

			tokenHash := s.sha256Hasher.Hash(token)

			req := domain.RegistrationRequest{
				FIO:         row.FIO,
				Email:       row.Email,
				PhoneNumber: row.PhoneNumber,
				Role:        row.Role,
				GroupID:     groupID,
				TokenHash:   tokenHash,
				ExpiresAt:   time.Now().Add(24 * time.Hour),
				Status:      enum.StatusPending,
			}

			if err := s.registrationRepository.Create(ctx, req); err != nil {
				return fmt.Errorf("create request: %w", err)
			}

			tasks = append(tasks, emailTask{
				email: row.Email,
				token: token,
			})
		}

		return nil
	})

	if err != nil {
		return err
	}

	for _, task := range tasks {
		err := s.emailSender.SendRegistrationEmail(
			ctx,
			task.email,
			"Complete your registration",
			fmt.Sprintf("Click link: http://127.0.0.1:5050/api/v1/register/complete?token=%s", task.token),
		)
		if err != nil {
			fmt.Printf("failed to send email: %v\n", err)
		}
	}

	return nil
}
