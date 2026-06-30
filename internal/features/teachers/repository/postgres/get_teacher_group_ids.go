package teachers_postgres_repository

import (
	"context"
	"fmt"
)

func (r *TeachersRepository) GetTeacherGroupIDs(
	ctx context.Context,
	teacherID int,
) ([]int, error) {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT group_id
	FROM myapp.teacher_groups
	WHERE teacher_id = $1;
	`

	db := r.dbFromContext(ctx)
	rows, err := db.Query(
		ctx,
		query,
		teacherID,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"query teacher groups: %w",
			err,
		)
	}
	defer rows.Close()

	groupIDs := make([]int, 0)

	for rows.Next() {

		var groupID int

		if err := rows.Scan(&groupID); err != nil {
			return nil, fmt.Errorf(
				"scan group id: %w",
				err,
			)
		}

		groupIDs = append(groupIDs, groupID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"iterate rows: %w",
			err,
		)
	}

	return groupIDs, nil
}
