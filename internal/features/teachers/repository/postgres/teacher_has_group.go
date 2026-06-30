package teachers_postgres_repository

import (
	"context"
	"fmt"
)

func (r *TeachersRepository) TeacherHasGroup(
	ctx context.Context,
	teacherID int,
	groupID int,
) (bool, error) {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT EXISTS
	(
		SELECT 1
		FROM myapp.teacher_groups
		WHERE teacher_id = $1
		  AND group_id = $2
	);
	`

	var exists bool
	db := r.dbFromContext(ctx)
	err := db.QueryRow(
		ctx,
		query,
		teacherID,
		groupID,
	).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf(
			"check teacher assignment: %w",
			err,
		)
	}

	return exists, nil
}
