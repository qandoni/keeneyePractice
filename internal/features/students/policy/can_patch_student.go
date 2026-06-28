package student_policy

import (
	"context"
	"fmt"

	core_auth "github.com/qandoni/keeneyePractice/internal/core/auth"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

func (p *StudentAccessPolicy) CanPatchStudent(
	ctx context.Context,
	auth core_auth.AuthInfo,
	studentID int,
) error {

	if auth.Role == enum.RoleAdmin {
		return nil
	}

	student, err := p.studentsRepository.GetStudent(
		ctx,
		studentID,
	)
	if err != nil {
		return fmt.Errorf("get student: %w", err)
	}

	switch auth.Role {

	case enum.RoleStudent:
		currentStudent, err := p.studentsRepository.GetStudentByUserID(
			ctx,
			auth.UserID,
		)
		if err != nil {
			return fmt.Errorf("get current student: %w", err)
		}

		if currentStudent.ID != student.ID {
			return core_errors.ErrAccessForbidden
		}

		return nil

	case enum.RoleTeacher:
		teacher, err := p.teachersRepository.GetTeacherByUserID(
			ctx,
			auth.UserID,
		)
		if err != nil {
			return fmt.Errorf("get teacher: %w", err)
		}

		ok, err := p.teachersRepository.TeacherHasGroup(
			ctx,
			teacher.ID,
			student.GroupID,
		)
		if err != nil {
			return fmt.Errorf("check teacher group: %w", err)
		}

		if !ok {
			return core_errors.ErrAccessForbidden
		}

		return nil

	default:
		return core_errors.ErrAccessForbidden
	}
}
