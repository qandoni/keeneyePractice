package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	core_postgres_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) PatchUser(
	ctx context.Context,
	id int,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE myapp.users
	SET
		login = $1,
		password_hash = $2,
		role = $3,
		version = version + 1
	WHERE id = $4
  	AND version = $5
	RETURNING
    	id,
    	version,
    	login,
    	password_hash,
    	role;
	`
	row := r.pool.QueryRow(
		ctx,
		query,
		user.Login,
		user.PasswordHash,
		user.Role,
		id,
		user.Version,
	)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Login,
		&userModel.PasswordHash,
		&userModel.Role,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"student with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}
	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.Login,
		userModel.PasswordHash,
		userModel.Role,
	)
	return userDomain, nil
}
