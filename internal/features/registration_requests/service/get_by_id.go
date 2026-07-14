package registration_service

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (s *RegistrationRequestService) GetByID(
	ctx context.Context,
	id int,
) (domain.RegistrationRequest, error) {
	request, err := s.registrationRepository.GetByID(ctx, id)
	if err != nil {
		return domain.RegistrationRequest{}, fmt.Errorf("get by id from repository: %w", err)
	}
	return request, nil
}
