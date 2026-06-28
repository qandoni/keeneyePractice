package teachers_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *TeachersService) GetTeacherByUserID(
	ctx context.Context,
	userID int,
) (domain.Teacher, error) {

	teacher, err := s.teachersRepository.GetTeacherByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return domain.Teacher{}, fmt.Errorf(
			"get teacher by user id: %w",
			err,
		)
	}

	return teacher, nil
}
