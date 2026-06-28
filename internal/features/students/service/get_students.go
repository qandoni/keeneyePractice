package students_service

import (
	"context"

	core_auth "github.com/qandoni/keeneyePractice/internal/core/auth"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

func (s *StudentsService) GetStudents(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Student, error) {

	auth, ok := core_auth.AuthInfoFromContext(ctx)
	if !ok {
		return nil, core_errors.ErrUnauthorized
	}

	scope, err := s.policy.CanGetStudents(ctx, auth)
	if err != nil {
		return nil, err
	}

	if scope.All {
		return s.studentsRepository.GetStudents(ctx, limit, offset)
	}

	if len(scope.GroupIDs) > 0 {
		return s.studentsRepository.GetStudentsByGroupIDs(
			ctx,
			scope.GroupIDs,
			limit,
			offset,
		)
	}

	return nil, core_errors.ErrAccessForbidden
}
