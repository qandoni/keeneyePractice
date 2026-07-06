package registration_service

import (
	"context"
	"fmt"
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
	registration_contracts "github.com/qandoni/keeneyePractice/internal/features/registration_requests/contracts"
	users_contracts "github.com/qandoni/keeneyePractice/internal/features/users/contracts"
)

func (s *RegistrationRequestService) Complete(
	ctx context.Context,
	input registration_contracts.CompleteInput,
) error {

	tokenHash := s.sha256Hasher.Hash(input.Token)
	return s.txManager.WithinTransaction(ctx, func(ctx context.Context) error {

		req, err := s.registrationRepository.GetByTokenHash(ctx, tokenHash)
		if err != nil {
			return fmt.Errorf("get registration request: %w", err)
		}

		if req.Status != enum.StatusPending {
			return fmt.Errorf("registration request already used")
		}

		if time.Now().After(req.ExpiresAt) {
			return fmt.Errorf("registration request expired")
		}

		createUserInput := users_contracts.CreateUserInput{
			Email:    req.Email,
			Password: input.Password,
			Role:     req.Role,
		}

		user, err := s.usersService.CreateUser(ctx, createUserInput)
		if err != nil {
			return fmt.Errorf("create user: %w", err)
		}

		switch req.Role {

		case enum.RoleStudent:

			if req.GroupID == nil {
				return fmt.Errorf("group_id is required for student")
			}

			student := domain.NewStudentUninitialized(
				user.ID,
				*req.GroupID,
				req.FIO,
				req.PhoneNumber,
			)

			if _, err := s.studentsService.CreateStudent(ctx, student); err != nil {
				return fmt.Errorf("create student: %w", err)
			}

		case enum.RoleTeacher:

			teacher := domain.NewTeacherUninitialized(
				user.ID,
				req.FIO,
				req.PhoneNumber,
			)

			if _, err := s.teachersService.CreateTeacher(ctx, teacher); err != nil {
				return fmt.Errorf("create teacher: %w", err)
			}

		default:
			return fmt.Errorf("invalid role: %s", req.Role)
		}

		if err := s.registrationRepository.UpdateStatus(
			ctx,
			req.ID,
			enum.StatusCompleted,
		); err != nil {
			return fmt.Errorf("update registration status: %w", err)
		}

		return nil
	})
}
