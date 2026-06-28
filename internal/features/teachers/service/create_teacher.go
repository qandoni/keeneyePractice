package teachers_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *TeachersService) CreateTeacher(
	ctx context.Context,
	teacher domain.Teacher,
) (domain.Teacher, error) {
	if err := teacher.Validate(); err != nil {
		return domain.Teacher{}, fmt.Errorf("validate student domain: %w", err)
	}
	teacher, err := s.teachersRepository.CreateTeacher(ctx, teacher)
	if err != nil {
		return domain.Teacher{}, fmt.Errorf("get teacher from repository: %w", err)
	}
	return teacher, nil
}
