package teachers_service

import (
	"context"
	"fmt"
)

func (s *TeachersService) RemoveFromGroup(
	ctx context.Context,
	teacherID int,
	groupID int,
) error {

	if err := s.teachersRepository.RemoveFromGroup(
		ctx,
		teacherID,
		groupID,
	); err != nil {
		return fmt.Errorf("remove teacher from group: %w", err)
	}

	return nil
}
