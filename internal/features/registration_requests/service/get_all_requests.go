package registration_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *RegistrationRequestService) GetAll(
	ctx context.Context,
) ([]domain.RegistrationRequest, error) {

	if err := s.registrationRepository.ExpireRequests(ctx); err != nil {
		return nil, fmt.Errorf("update expired requests: %w", err)
	}

	requests, err := s.registrationRepository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("get registration requests: %w", err)
	}

	return requests, nil
}
