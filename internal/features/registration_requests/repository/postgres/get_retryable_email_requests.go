package registration_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *RegistrationRequestsRepository) GetRetryableEmailRequests(
	ctx context.Context,
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
	WHERE
    status = 'pending'
    AND email_status IN ('pending', 'failed')
    AND (
        last_email_attempt_at IS NULL
        OR last_email_attempt_at <= NOW() - INTERVAL '5 minutes'
    )
	ORDER BY id;
	`
	db := r.dbFromContext(ctx)
	rows, err := db.Query(
		ctx,
		query,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"query retryable requests: %w",
			err,
		)
	}

	defer rows.Close()

	var requests []domain.RegistrationRequest

	for rows.Next() {

		var model RegistrationRequestModel

		err := rows.Scan(
			&model.ID,
			&model.Version,
			&model.FIO,
			&model.Email,
			&model.PhoneNumber,
			&model.Role,
			&model.GroupID,
			&model.TokenHash,
			&model.Status,
			&model.EmailStatus,
			&model.EmailRetryCount,
			&model.LastEmailAttempt,
			&model.EmailSentAt,
			&model.ExpiresAt,
			&model.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf(
				"scan request: %w",
				err,
			)
		}

		requests = append(
			requests,
			modelToDomain(model),
		)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"iterate requests: %w",
			err,
		)
	}

	return requests, nil
}
