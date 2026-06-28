package teachers_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

func (s *TeachersService) GetTeachers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Teacher, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	teachers, err := s.teachersRepository.GetTeachers(ctx, limit, offset)
	if err != nil {
		return []domain.Teacher{}, fmt.Errorf("get students from repository: %w", err)
	}
	return teachers, nil
}
