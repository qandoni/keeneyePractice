package teachers_service

import (
	"context"
	"fmt"
)

func (s *TeachersService) GetTeacherGroupIDs(
	ctx context.Context,
	teacherID int,
) ([]int, error) {

	groupIDs, err := s.teachersRepository.GetTeacherGroupIDs(
		ctx,
		teacherID,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"get teacher group ids: %w",
			err,
		)
	}

	return groupIDs, nil
}
