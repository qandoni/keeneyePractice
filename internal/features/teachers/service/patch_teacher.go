package teachers_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *TeachersService) PatchTeacher(
	ctx context.Context,
	id int,
	patch domain.TeacherPatch,
) (domain.Teacher, error) {
	teacher, err := s.teachersRepository.GetTeacher(ctx, id)
	if err != nil {
		return domain.Teacher{}, fmt.Errorf("get teacher: %w", err)
	}
	if err := teacher.ApplyPatch(patch); err != nil {
		return domain.Teacher{}, fmt.Errorf("apply teacher patch: %w", err)
	}
	patchedTeacher, err := s.teachersRepository.PatchTeacher(ctx, id, teacher)
	if err != nil {
		return domain.Teacher{}, fmt.Errorf("patch teacher: %w", err)
	}
	return patchedTeacher, nil
}
