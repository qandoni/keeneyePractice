package teachers_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *TeachersService) GetTeacher(
	ctx context.Context,
	id int,
) (domain.Teacher, error) {
	teacher, err := s.teachersRepository.GetTeacher(ctx, id)
	if err != nil {
		return domain.Teacher{}, fmt.Errorf("get teacher from repository: %w", err)
	}
	return teacher, nil
}
