package students_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *StudentsService) GetStudent(
	ctx context.Context,
	id int,
) (domain.Student, error) {
	student, err := s.studentsRepository.GetStudent(ctx, id)
	if err != nil {
		return domain.Student{}, fmt.Errorf("get user from repository: %w", err)
	}
	return student, nil
}
