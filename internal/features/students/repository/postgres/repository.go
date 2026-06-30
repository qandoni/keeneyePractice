package students_postgres_repository

import (
	"context"
	"time"

	core_postgres "github.com/qandoni/keeneyePractice/internal/core/repository/postgres"
)

type StudentsRepository struct {
	db      core_postgres.DB
	timeout time.Duration
}

func NewStudentsRepository(
	db core_postgres.DB,
	timeout time.Duration,
) *StudentsRepository {
	return &StudentsRepository{
		db:      db,
		timeout: timeout,
	}
}

func (r *StudentsRepository) dbFromContext(
	ctx context.Context,
) core_postgres.DB {

	db := core_postgres.DBFromContext(ctx)
	if db != nil {
		return db
	}

	return r.db
}
