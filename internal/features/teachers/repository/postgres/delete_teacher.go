package teachers_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

func (r *TeachersRepository) DeleteTeacher(
	ctx context.Context,
	id int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM myapp.teachers
	WHERE id=$1;
	`
	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("teacher with id='%d': %w", id, core_errors.ErrNotFound)
	}
	return nil
}
