package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	core_postgres_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT id, version, email, password_hash, role
	FROM myapp.users
	WHERE email=$1;
	`

	db := r.dbFromContext(ctx)
	row := db.QueryRow(ctx, query, email)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Email,
		&userModel.PasswordHash,
		&userModel.Role,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with email='%v': %w", core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.Email,
		userModel.PasswordHash,
		userModel.Role,
	)
	return userDomain, nil
}
