package registration_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *RegistrationRequestsRepository) GetByTokenHash(
	ctx context.Context,
	hash string,
) (domain.RegistrationRequest, error) {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT
		id,
		version,
		fio,
		email,
		phone_number,
		role,
		group_id,
		token_hash,
		expires_at,
		status
	FROM myapp.registration_requests
	WHERE token_hash = $1
	`

	row := r.db.QueryRow(ctx, query, hash)

	var m RegistrationRequestModel

	err := row.Scan(
		&m.ID,
		&m.Version,
		&m.FIO,
		&m.Email,
		&m.PhoneNumber,
		&m.Role,
		&m.GroupID,
		&m.TokenHash,
		&m.ExpiresAt,
		&m.Status,
	)

	if err != nil {
		return domain.RegistrationRequest{}, fmt.Errorf(
			"get by token hash: %w",
			err,
		)
	}

	return modelToDomain(m), nil
}
