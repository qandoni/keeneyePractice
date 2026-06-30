package groups_postgres_repository

import (
	"context"
	"time"

	core_postgres "github.com/qandoni/keeneyePractice/internal/core/repository/postgres"
)

type GroupsRepository struct {
	db      core_postgres.DB
	timeout time.Duration
}

func NewGroupsRepository(
	db core_postgres.DB,
	timeout time.Duration,
) *GroupsRepository {
	return &GroupsRepository{
		db:      db,
		timeout: timeout,
	}
}

func (r *GroupsRepository) dbFromContext(
	ctx context.Context,
) core_postgres.DB {

	db := core_postgres.DBFromContext(ctx)
	if db != nil {
		return db
	}

	return r.db
}
