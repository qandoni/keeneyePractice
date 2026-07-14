package registration_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
)

func (r *RegistrationRequestsRepository) Create(
	ctx context.Context,
	req domain.RegistrationRequest,
) (domain.RegistrationRequest, error) {
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
	RETURNING id, version, fio, email, phone_number, role, group_id, token_hash, status, email_status, email_retry_count, last_email_attempt_at, email_sent_at, expires_at, created_at
	`
	db := r.dbFromContext(ctx)
	row := db.QueryRow(ctx, query,
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
		return domain.RegistrationRequest{}, fmt.Errorf("create registration request: %w", err)
	}

	registrationRequest := domain.NewRegistrationRequest(
		registrationRequestModel.ID,
		registrationRequestModel.Version,
		registrationRequestModel.FIO,
		registrationRequestModel.Email,
		registrationRequestModel.PhoneNumber,
		enum.Role(registrationRequestModel.Role),
		registrationRequestModel.GroupID,
		registrationRequestModel.TokenHash,
		enum.RegistrationStatus(registrationRequestModel.Status),
		enum.EmailStatus(registrationRequestModel.EmailStatus),
		registrationRequestModel.EmailRetryCount,
		registrationRequestModel.LastEmailAttempt,
		registrationRequestModel.EmailSentAt,
		registrationRequestModel.ExpiresAt,
		registrationRequestModel.CreatedAt,
	)

	return registrationRequest, nil
}
