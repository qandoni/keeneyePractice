package admin_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
	admin_contracts "github.com/qandoni/keeneyePractice/internal/features/admin/contracts"
	users_contracts "github.com/qandoni/keeneyePractice/internal/features/users/contracts"
)

func (s *AdminService) CreateUser(
	ctx context.Context,
	cmd admin_contracts.CreateUserCommand,
) (domain.User, error) {

	var createdUser domain.User

	err := s.txManager.WithinTransaction(ctx, func(ctx context.Context) error {

		input := users_contracts.CreateUserInput{
			Login:    cmd.Login,
			Password: cmd.Password,
			Role:     cmd.Role,
		}

		user, err := s.usersService.CreateUser(ctx, input)
		if err != nil {
			return fmt.Errorf("create user: %w", err)
		}

		createdUser = user

		switch cmd.Role {

		case enum.RoleStudent:

			if cmd.GroupID == nil {
				return fmt.Errorf("group_id is required for student")
			}

			student := domain.NewStudentUninitialized(
				user.ID,
				*cmd.GroupID,
				cmd.FIO,
				cmd.PhoneNumber,
			)

			if _, err := s.studentsService.CreateStudent(ctx, student); err != nil {
				return fmt.Errorf("create student: %w", err)
			}

		case enum.RoleTeacher:

			teacher := domain.NewTeacherUninitialized(
				user.ID,
				cmd.FIO,
				cmd.PhoneNumber,
			)

			if _, err := s.teachersService.CreateTeacher(ctx, teacher); err != nil {
				return fmt.Errorf("create teacher: %w", err)
			}

		case enum.RoleAdmin:

		default:
			return fmt.Errorf("unknown role: %s", cmd.Role)
		}

		return nil
	})

	if err != nil {
		return domain.User{}, err
	}

	return createdUser, nil
}
