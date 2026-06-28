package students_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *StudentsService) GetStudentByUserID(
	ctx context.Context,
	userID int,
) (domain.Student, error) {

	student, err := s.studentsRepository.GetStudentByUserID(ctx, userID)
	if err != nil {
		return domain.Student{}, fmt.Errorf("get student by user id: %w", err)
	}

	return student, nil
}
