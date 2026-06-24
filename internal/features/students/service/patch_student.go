package students_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *StudentsService) PatchStudent(
	ctx context.Context,
	id int,
	patch domain.StudentPatch,
) (domain.Student, error) {
	student, err := s.studentsRepository.GetStudent(ctx, id)
	if err != nil {
		return domain.Student{}, fmt.Errorf("get student: %w", err)
	}
	if err := student.ApplyPatch(patch); err != nil {
		return domain.Student{}, fmt.Errorf("apply student patch: %w", err)
	}
	patchedStudent, err := s.studentsRepository.PatchStudent(ctx, id, student)
	if err != nil {
		return domain.Student{}, fmt.Errorf("patch student: %w", err)
	}
	return patchedStudent, nil
}
