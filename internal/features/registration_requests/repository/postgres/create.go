package registration_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *RegistrationRequestsRepository) Create(
	ctx context.Context,
	req domain.RegistrationRequest,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	INSERT INTO myapp.registration_requests(
		fio,
		email,
		phone_number,
		role,
		group_id,
		token_hash,
		expires_at,
		status	
	)
	VALUES(
		$1, $2, $3, $4, $5, $6, $7, $8
	)
	`

	_, err := r.db.Exec(ctx, query,
		req.FIO,
		req.Email,
		req.PhoneNumber,
		req.Role,
		req.GroupID,
		req.TokenHash,
		req.ExpiresAt,
		req.Status,
	)
	if err != nil {
		return fmt.Errorf("create registration request: %w", err)
	}
	return nil
}
