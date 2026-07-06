package registration_postgres_repository

import (
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
