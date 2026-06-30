package groups_postgres_repository

import (
	"context"
	"fmt"
)

func (r *GroupsRepository) DeleteGroup(
	ctx context.Context,
	id int,
) error {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	DELETE FROM myapp.groups
	WHERE id = $1
	`
	db := r.dbFromContext(ctx)
	commandTag, err := db.Exec(
		ctx,
		query,
		id,
	)
	if err != nil {
		return fmt.Errorf("delete group: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("group not found")
	}

	return nil
}
