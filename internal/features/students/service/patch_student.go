package students_service

import (
	"context"
	"fmt"

	core_auth "github.com/qandoni/keeneyePractice/internal/core/auth"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

func (s *StudentsService) PatchStudent(
	ctx context.Context,
	studentID int,
	patch domain.StudentPatch,
) (domain.Student, error) {

	auth, ok := core_auth.AuthInfoFromContext(ctx)
	if !ok {
		return domain.Student{}, core_errors.ErrUnauthorized
	}

	if err := s.policy.CanPatchStudent(ctx, auth, studentID); err != nil {
		return domain.Student{}, err
	}

	student, err := s.studentsRepository.GetStudent(ctx, studentID)
	if err != nil {
		return domain.Student{}, fmt.Errorf("get student: %w", err)
	}

	if err := student.ApplyPatch(patch); err != nil {
		return domain.Student{}, err
	}

	updated, err := s.studentsRepository.PatchStudent(ctx, studentID, student)
	if err != nil {
		return domain.Student{}, fmt.Errorf("update student: %w", err)
	}

	return updated, nil
}
