package students_postgres_repository

import core_postgres_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool"

type StudentsRepository struct {
	pool core_postgres_pool.Pool
}

func NewStudentsRepository(
	pool core_postgres_pool.Pool,
) *StudentsRepository {
	return &StudentsRepository{
		pool: pool,
	}
}
