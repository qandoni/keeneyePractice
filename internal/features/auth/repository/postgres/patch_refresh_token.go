package auth_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	core_postgres_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool"
)

func (r *RefreshTokensRepository) PatchRefreshToken(
	ctx context.Context,
	token domain.RefreshToken,
) (domain.RefreshToken, error) {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	UPDATE myapp.refresh_tokens
	SET
		token_hash = $1,
		expires_at = $2,
		version = version + 1
	WHERE
		id = $3
	AND version = $4
	RETURNING
		id,
		version,
		user_id,
		token_hash,
		expires_at;
	`

	row := r.db.QueryRow(
		ctx,
		query,
		token.TokenHash,
		token.ExpiresAt,
		token.ID,
		token.Version,
	)

	var model RefreshTokenModel

	err := row.Scan(
		&model.ID,
		&model.Version,
		&model.UserID,
		&model.TokenHash,
		&model.ExpiresAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.RefreshToken{}, fmt.Errorf(
				"refresh token concurrently modified: %w",
				core_errors.ErrConflict,
			)
		}

		return domain.RefreshToken{}, fmt.Errorf(
			"scan error: %w",
			err,
		)
	}

	return domain.NewRefreshToken(
		model.ID,
		model.Version,
		model.UserID,
		model.TokenHash,
		model.ExpiresAt,
	), nil
}
