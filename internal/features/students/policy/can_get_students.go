package student_policy

import (
	"context"
	"fmt"

	core_auth "github.com/qandoni/keeneyePractice/internal/core/auth"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

func (p *StudentAccessPolicy) CanGetStudents(
	ctx context.Context,
	auth core_auth.AuthInfo,
) (StudentScope, error) {

	switch auth.Role {

	case enum.RoleAdmin:
		return StudentScope{
			All: true,
		}, nil

	case enum.RoleStudent:
		student, err := p.studentsRepository.GetStudentByUserID(
			ctx,
			auth.UserID,
		)
		if err != nil {
			return StudentScope{}, fmt.Errorf(
				"get student: %w",
				err,
			)
		}

		return StudentScope{
			GroupIDs: []int{
				student.GroupID,
			},
		}, nil

	case enum.RoleTeacher:
		teacher, err := p.teachersRepository.GetTeacherByUserID(
			ctx,
			auth.UserID,
		)
		if err != nil {
			return StudentScope{}, fmt.Errorf(
				"get teacher: %w",
				err,
			)
		}

		groupIDs, err := p.teachersRepository.GetTeacherGroupIDs(
			ctx,
			teacher.ID,
		)
		if err != nil {
			return StudentScope{}, fmt.Errorf(
				"get teacher groups: %w",
				err,
			)
		}

		return StudentScope{
			GroupIDs: groupIDs,
		}, nil

	default:
		return StudentScope{}, core_errors.ErrAccessForbidden
	}
}
