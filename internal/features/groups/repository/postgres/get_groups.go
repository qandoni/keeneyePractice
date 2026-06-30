package groups_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *GroupsRepository) GetGroups(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Group, error) {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT
		id,
		version,
		name
	FROM myapp.groups
	ORDER BY id
	LIMIT $1 OFFSET $2
	`
	db := r.dbFromContext(ctx)
	rows, err := db.Query(
		ctx,
		query,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("query groups: %w", err)
	}
	defer rows.Close()

	var groups []domain.Group

	for rows.Next() {
		var model GroupModel

		if err := rows.Scan(
			&model.ID,
			&model.Version,
			&model.Name,
		); err != nil {
			return nil, fmt.Errorf("scan group: %w", err)
		}

		groups = append(
			groups,
			domain.NewGroup(
				model.ID,
				model.Version,
				model.Name,
			),
		)
	}

	return groups, nil
}
