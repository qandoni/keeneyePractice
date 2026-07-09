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
		status,
		email_status,
		email_retry_count,
		last_email_attempt_at,
		email_sent_at,
		expires_at
	)
	VALUES(
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
	)
	`
	db := r.dbFromContext(ctx)
	_, err := db.Exec(ctx, query,
		req.FIO,
		req.Email,
		req.PhoneNumber,
		req.Role,
		req.GroupID,
		req.TokenHash,
		req.Status,
		req.EmailStatus,
		req.EmailRetryCount,
		req.LastEmailAttempt,
		req.EmailSentAt,
		req.ExpiresAt,
	)
	if err != nil {
		return fmt.Errorf("create registration request: %w", err)
	}
	return nil
}
