package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO myapp.users(login, password_hash, role)
	VALUES($1, $2, $3)
	RETURNING id, version, login, password_hash, role
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.Login,
		user.PasswordHash,
		user.Role,
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
