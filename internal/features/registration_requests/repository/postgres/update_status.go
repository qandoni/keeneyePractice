package registration_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/enum"
)

func (r *RegistrationRequestsRepository) UpdateStatus(
	ctx context.Context,
	id int,
	status enum.RegistrationStatus,
) error {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
		UPDATE myapp.registration_requests
		SET status = $1,
		    version = version + 1
		WHERE id = $2
	`
	db := r.dbFromContext(ctx)
	result, err := db.Exec(ctx, query, status, id)
	if err != nil {
		return fmt.Errorf("update registration status: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("registration request not found")
	}

	return nil
}
