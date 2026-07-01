package auth_postgres_repository

import (
	"time"

	core_postgres "github.com/qandoni/keeneyePractice/internal/core/repository/postgres"
)

type RefreshTokensRepository struct {
	db      core_postgres.DB
	timeout time.Duration
}

func NewRefreshTokensRepository(
	db core_postgres.DB,
	timeout time.Duration,
) *RefreshTokensRepository {

	return &RefreshTokensRepository{
		db:      db,
		timeout: timeout,
	}
}
