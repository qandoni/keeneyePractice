package teachers_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *TeachersRepository) AssignToGroup(
	ctx context.Context,
	assignment domain.TeacherGroupAssignment,
) error {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	INSERT INTO myapp.teacher_groups
	(
		teacher_id,
		group_id
	)
	VALUES
	(
		$1,
		$2
	);
	`

	db := r.dbFromContext(ctx)
	_, err := db.Exec(
		ctx,
		query,
		assignment.TeacherID,
		assignment.GroupID,
	)
	if err != nil {
		return fmt.Errorf(
			"insert teacher assignment: %w",
			err,
		)
	}

	return nil
}
