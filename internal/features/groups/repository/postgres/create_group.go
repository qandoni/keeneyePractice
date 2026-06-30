package groups_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *GroupsRepository) CreateGroup(
	ctx context.Context,
	group domain.Group,
) (domain.Group, error) {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	INSERT INTO myapp.groups(name)
	VALUES($1)
	RETURNING id, version, name
	`
	db := r.dbFromContext(ctx)
	row := db.QueryRow(
		ctx,
		query,
		group.Name,
	)

	var model GroupModel

	err := row.Scan(
		&model.ID,
		&model.Version,
		&model.Name,
	)
	if err != nil {
		return domain.Group{}, fmt.Errorf("scan group: %w", err)
	}

	return domain.NewGroup(
		model.ID,
		model.Version,
		model.Name,
	), nil
}
