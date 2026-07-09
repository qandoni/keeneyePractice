package registration_postgres_repository

import (
	"context"
	"fmt"

	registration_contracts "github.com/qandoni/keeneyePractice/internal/features/registration_requests/contracts"
)

func (r *RegistrationRequestsRepository) UpdateAfterEmailAttempt(
	ctx context.Context,
	input registration_contracts.UpdateEmailAttemptInput,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	UPDATE myapp.registration_requests
	SET
		token_hash = $1,
		expires_at = $2,
		email_status = $3,
		email_retry_count = $4,
		last_email_attempt_at = $5,
		email_sent_at = $6,
		version = version + 1
	WHERE id = $7
	AND version = $8
	`

	db := r.dbFromContext(ctx)
	result, err := db.Exec(
		ctx,
		query,
		input.TokenHash,
		input.ExpiresAt,
		input.EmailStatus,
		input.EmailRetryCount,
		input.LastEmailAttempt,
		input.EmailSentAt,
		input.ID,
		input.Version,
	)

	if err != nil {
		return fmt.Errorf("apply update in repository: %w", err)
	}

	rows := result.RowsAffected()

	if rows == 0 {
		return fmt.Errorf(
			"registration request not updated: version conflict",
		)
	}

	return nil
}
