package teachers_service

import (
	"context"
	"fmt"
)

func (s *TeachersService) DeleteTeacher(
	ctx context.Context,
	id int,
) error {
	if err := s.teachersRepository.DeleteTeacher(ctx, id); err != nil {
		return fmt.Errorf("delete teacher: %w", err)
	}
	return nil
}
