package students_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *StudentsService) CreateStudent(
	ctx context.Context,
	student domain.Student,
) (domain.Student, error) {
	if err := student.Validate(); err != nil {
		return domain.Student{}, fmt.Errorf("validate student domain: %w", err)
	}
	student, err := s.studentsRepository.CreateStudent(ctx, student)
	if err != nil {
		return domain.Student{}, fmt.Errorf("create user: %w", err)
	}
	return student, nil
}
