package auth_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *RefreshTokensRepository) Save(
	ctx context.Context,
	token domain.RefreshToken,
) error {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	INSERT INTO myapp.refresh_tokens
	(
		version,
		token_hash,
		user_id,
		expires_at
	)
	VALUES
	(
		$1,
		$2,
		$3,
		$4
	)
	ON CONFLICT (user_id)
	DO UPDATE SET
		token_hash = EXCLUDED.token_hash,
		expires_at = EXCLUDED.expires_at,
		version = refresh_tokens.version + 1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		token.Version,
		token.TokenHash,
		token.UserID,
		token.ExpiresAt,
	)

	if err != nil {
		return fmt.Errorf("create refresh token: %w", err)
	}

	return nil
}
