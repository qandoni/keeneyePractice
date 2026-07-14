package registration_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *RegistrationRequestsRepository) GetByID(
	ctx context.Context,
	id int,
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
	WHERE id = $1
	`

	db := r.dbFromContext(ctx)
	row := db.QueryRow(ctx, query, id)

	var registrationRequestModel RegistrationRequestModel

	err := row.Scan(
		&registrationRequestModel.ID,
		&registrationRequestModel.Version,
		&registrationRequestModel.FIO,
		&registrationRequestModel.Email,
		&registrationRequestModel.PhoneNumber,
		&registrationRequestModel.Role,
		&registrationRequestModel.GroupID,
		&registrationRequestModel.TokenHash,
		&registrationRequestModel.Status,
		&registrationRequestModel.EmailStatus,
		&registrationRequestModel.EmailRetryCount,
		&registrationRequestModel.LastEmailAttempt,
		&registrationRequestModel.EmailSentAt,
		&registrationRequestModel.ExpiresAt,
		&registrationRequestModel.CreatedAt,
	)
	if err != nil {
		return domain.RegistrationRequest{}, fmt.Errorf(
			"get by request id: %w", err,
		)
	}
	return modelToDomain(registrationRequestModel), nil
}
