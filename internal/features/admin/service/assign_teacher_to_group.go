package admin_service

import "context"

func (s *AdminService) AssignTeacherToGroup(
	ctx context.Context,
	teacherID int,
	groupID int,
) error {

	return s.teachersService.AssignToGroup(ctx, teacherID, groupID)
}
