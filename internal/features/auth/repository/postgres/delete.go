package auth_postgres_repository

import (
	"context"
	"fmt"
)

func (r *RefreshTokensRepository) Delete(
	ctx context.Context,
	tokenHash string,
) error {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	DELETE FROM myapp.refresh_tokens
	WHERE token_hash=$1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		tokenHash,
	)

	if err != nil {
		return fmt.Errorf(
			"delete refresh token: %w",
			err,
		)
	}

	return nil
}
