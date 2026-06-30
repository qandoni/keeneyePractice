package groups_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *GroupsRepository) PatchGroup(
	ctx context.Context,
	group domain.Group,
) (domain.Group, error) {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	UPDATE myapp.groups
	SET
		name = $1,
		version = version + 1
	WHERE
		id = $2
		AND version = $3
	RETURNING
		id,
		version,
		name
	`
	db := r.dbFromContext(ctx)
	row := db.QueryRow(
		ctx,
		query,
		group.Name,
		group.ID,
		group.Version,
	)

	var model GroupModel

	err := row.Scan(
		&model.ID,
		&model.Version,
		&model.Name,
	)
	if err != nil {
		return domain.Group{}, fmt.Errorf("patch group: %w", err)
	}

	return domain.NewGroup(
		model.ID,
		model.Version,
		model.Name,
	), nil
}
