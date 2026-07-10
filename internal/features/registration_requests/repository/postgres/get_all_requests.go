package registration_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *RegistrationRequestsRepository) GetAll(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.RegistrationRequest, error) {

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
	ORDER BY id
	LIMIT $1
	OFFSET $2
	`
	db := r.dbFromContext(ctx)
	rows, err := db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query requests: %w", err)
	}
	defer rows.Close()

	requests := make([]domain.RegistrationRequest, 0)

	for rows.Next() {
		var m RegistrationRequestModel

		err := rows.Scan(
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
			return nil, fmt.Errorf("scan request: %w", err)
		}

		requests = append(requests, modelToDomain(m))
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate rows: %w", err)
	}

	return requests, nil
}
