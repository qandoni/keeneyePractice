package teachers_service

import "context"

func (s *TeachersService) TeacherHasGroup(
	ctx context.Context,
	teacherID int,
	groupID int,
) (bool, error) {

	groupIDs, err := s.GetTeacherGroupIDs(
		ctx,
		teacherID,
	)
	if err != nil {
		return false, err
	}

	for _, id := range groupIDs {
		if id == groupID {
			return true, nil
		}
	}

	return false, nil
}
