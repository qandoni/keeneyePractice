package groups_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

func (r *GroupsRepository) GetGroup(
	ctx context.Context,
	id int,
) (domain.Group, error) {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT
		id,
		version,
		name
	FROM myapp.groups
	WHERE id = $1
	`
	db := r.dbFromContext(ctx)
	row := db.QueryRow(
		ctx,
		query,
		id,
	)

	var model GroupModel

	err := row.Scan(
		&model.ID,
		&model.Version,
		&model.Name,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return domain.Group{}, core_errors.ErrNotFound
	}

	if err != nil {
		return domain.Group{}, fmt.Errorf("scan group: %w", err)
	}

	return domain.NewGroup(
		model.ID,
		model.Version,
		model.Name,
	), nil
}
