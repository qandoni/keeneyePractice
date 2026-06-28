package students_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *StudentsService) GetStudentsByGroupIDs(
	ctx context.Context,
	groupIDs []int,
	limit *int,
	offset *int,
) ([]domain.Student, error) {
	students, err := s.studentsRepository.GetStudentsByGroupIDs(
		ctx,
		groupIDs,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("get students by group ids: %w", err)
	}

	return students, nil
}
