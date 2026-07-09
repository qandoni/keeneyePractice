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
		status,
		email_status,
		email_retry_count,
		last_email_attempt_at,
		email_sent_at,
		expires_at,
		created_at
	FROM myapp.registration_requests
	WHERE token_hash = $1
	`
	db := r.dbFromContext(ctx)
	row := db.QueryRow(ctx, query, hash)

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
		&m.Status,
		&m.EmailStatus,
		&m.EmailRetryCount,
		&m.LastEmailAttempt,
		&m.EmailSentAt,
		&m.ExpiresAt,
		&m.CreatedAt,
	)

	if err != nil {
		return domain.RegistrationRequest{}, fmt.Errorf(
			"get by token hash: %w",
			err,
		)
	}

	return modelToDomain(m), nil
}
