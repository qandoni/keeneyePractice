package auth_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *RefreshTokensRepository) Get(
	ctx context.Context,
	tokenHash string,
) (domain.RefreshToken, error) {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT
		id,
		version,
		token_hash,
		user_id,
		expires_at
	FROM myapp.refresh_tokens
	WHERE token_hash=$1
	`

	row := r.db.QueryRow(
		ctx,
		query,
		tokenHash,
	)

	var model RefreshTokenModel

	err := row.Scan(
		&model.ID,
		&model.Version,
		&model.TokenHash,
		&model.UserID,
		&model.ExpiresAt,
	)

	if err != nil {
		return domain.RefreshToken{}, fmt.Errorf(
			"scan refresh token: %w",
			err,
		)
	}

	return modelToDomain(model), nil
}
