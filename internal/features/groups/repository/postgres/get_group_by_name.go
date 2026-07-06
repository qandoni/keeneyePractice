package groups_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	core_postgres_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool"
)

func (r *GroupsRepository) GetGroupByName(
	ctx context.Context,
	name string,
) (domain.Group, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT 
		id,
		version,
		name
	FROM myapp.groups
	WHERE name = $1
	`

	db := r.dbFromContext(ctx)
	row := db.QueryRow(ctx, query, name)

	var m GroupModel

	err := row.Scan(
		&m.ID,
		&m.Version,
		&m.Name,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Group{}, fmt.Errorf(
				"group by name not found: %w",
				core_errors.ErrNotFound,
			)
		}
		return domain.Group{}, fmt.Errorf("scan error: %w", err)
	}
	return domain.NewGroup(
		m.ID,
		m.Version,
		m.Name,
	), nil
}
