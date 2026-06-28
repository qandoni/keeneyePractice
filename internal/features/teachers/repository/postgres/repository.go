package teachers_postgres_repository

import core_postgres_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool"

type TeachersRepository struct {
	pool core_postgres_pool.Pool
}

func NewUsersRepository(
	pool core_postgres_pool.Pool,
) *TeachersRepository {
	return &TeachersRepository{
		pool: pool,
	}
}
