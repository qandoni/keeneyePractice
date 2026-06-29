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
	input := users_contracts.CreateUserInput{
		Login:    cmd.Login,
		Password: cmd.Password,
		Role:     cmd.Role,
	}
	user, err := s.usersService.CreateUser(ctx, input)
	fmt.Printf("created user: %+v\n", user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	switch cmd.Role {
	case enum.RoleStudent:
		if cmd.GroupID == nil {
			return domain.User{}, fmt.Errorf("group_id is required for student")
		}
		student := domain.NewStudentUninitialized(
			user.ID,
			*cmd.GroupID,
			cmd.FIO,
			cmd.PhoneNumber,
		)
		_, err := s.studentsService.CreateStudent(ctx, student)
		if err != nil {
			return domain.User{}, fmt.Errorf("create student: %w", err)
		}
	case enum.RoleTeacher:
		teacher := domain.NewTeacherUninitialized(user.ID, cmd.FIO, cmd.PhoneNumber)
		_, err := s.teachersService.CreateTeacher(ctx, teacher)
		if err != nil {
			return domain.User{}, fmt.Errorf("create teacher: %w", err)
		}
	case enum.RoleAdmin:
	default:
		return domain.User{}, fmt.Errorf(
			"unknown role: %s", cmd.Role,
		)
	}
	return user, nil

}
