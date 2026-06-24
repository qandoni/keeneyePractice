package students_service

import (
	"context"
	"fmt"
)

func (s *StudentsService) DeleteStudent(
	ctx context.Context,
	id int,
) error {
	if err := s.studentsRepository.DeleteStudent(ctx, id); err != nil {
		return fmt.Errorf("delete student: %w", err)
	}
	return nil
}
