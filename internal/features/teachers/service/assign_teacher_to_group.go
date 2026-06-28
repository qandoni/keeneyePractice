package teachers_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *TeachersService) AssignToGroup(
	ctx context.Context,
	teacherID int,
	groupID int,
) error {
	assignment := domain.TeacherGroupAssignment{
		TeacherID: teacherID,
		GroupID:   groupID,
	}
	if err := s.teachersRepository.AssignToGroup(
		ctx,
		assignment,
	); err != nil {
		return fmt.Errorf("assign teacher to group: %w", err)
	}
	return nil
}
