package teachers_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

func (r *TeachersRepository) RemoveFromGroup(
	ctx context.Context,
	teacherID int,
	groupID int,
) error {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	DELETE
	FROM myapp.teacher_groups
	WHERE teacher_id = $1
	  AND group_id = $2;
	`

	db := r.dbFromContext(ctx)
	commandTag, err := db.Exec(
		ctx,
		query,
		teacherID,
		groupID,
	)
	if err != nil {
		return fmt.Errorf(
			"delete teacher assignment: %w",
			err,
		)
	}

	if commandTag.RowsAffected() == 0 {
		return core_errors.ErrNotFound
	}

	return nil
}
