package registration_postgres_repository

import (
	"context"
	"fmt"
)

func (r *RegistrationRequestsRepository) ExpireRequests(
	ctx context.Context,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
		UPDATE myapp.registration_requests
		SET
			status = 'expired',
			version = version + 1
		WHERE
			status = 'pending'
			AND expires_at <= NOW()
	`

	_, err := r.db.Exec(ctx, query)
	fmt.Println("changed status somewhere...")
	if err != nil {
		return fmt.Errorf("expire registration requests: %w", err)
	}

	return nil
}
