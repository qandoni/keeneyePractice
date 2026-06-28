package admin_service

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *AdminService) GetTeachers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Teacher, error) {

	return s.teachersService.GetTeachers(ctx, limit, offset)
}
