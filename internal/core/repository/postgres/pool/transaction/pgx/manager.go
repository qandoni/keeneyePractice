package core_postgres_pool_transaction

import "github.com/jackc/pgx/v5/pgxpool"

type Manager struct {
	pool *pgxpool.Pool
}

func NewManager(pool *pgxpool.Pool) *Manager {
	return &Manager{
		pool: pool,
	}
}
