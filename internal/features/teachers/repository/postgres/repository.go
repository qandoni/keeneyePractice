package teachers_postgres_repository

import (
	"context"
	"time"

	core_postgres "github.com/qandoni/keeneyePractice/internal/core/repository/postgres"
)

type TeachersRepository struct {
	db      core_postgres.DB
	timeout time.Duration
}

func NewUsersRepository(
	db core_postgres.DB,
	timeout time.Duration,
) *TeachersRepository {
	return &TeachersRepository{
		db:      db,
		timeout: timeout,
	}
}

func (r *TeachersRepository) dbFromContext(
	ctx context.Context,
) core_postgres.DB {

	db := core_postgres.DBFromContext(ctx)
	if db != nil {
		return db
	}

	return r.db
}
