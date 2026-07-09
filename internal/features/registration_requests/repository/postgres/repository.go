package registration_postgres_repository

import (
	"context"
	"time"

	core_postgres "github.com/qandoni/keeneyePractice/internal/core/repository/postgres"
)

type RegistrationRequestsRepository struct {
	db      core_postgres.DB
	timeout time.Duration
}

func NewRegistrationRequestsRepository(
	db core_postgres.DB,
	timeout time.Duration,
) *RegistrationRequestsRepository {
	return &RegistrationRequestsRepository{
		db:      db,
		timeout: timeout,
	}
}

func (r *RegistrationRequestsRepository) dbFromContext(
	ctx context.Context,
) core_postgres.DB {

	db := core_postgres.DBFromContext(ctx)
	if db != nil {
		return db
	}
	return r.db
}
