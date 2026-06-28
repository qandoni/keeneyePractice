package admin_service

import "context"

func (s *AdminService) RemoveTeacherFromGroup(
	ctx context.Context,
	teacherID int,
	groupID int,
) error {

	return s.teachersService.RemoveFromGroup(ctx, teacherID, groupID)
}
