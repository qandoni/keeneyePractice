package students_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	core_postgres_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool"
)

func (r *StudentsRepository) GetStudentByUserID(
	ctx context.Context,
	userID int,
) (domain.Student, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	SELECT
		id,
		version,
		user_id,
		group_id,
		fio,
		phone_number
	FROM myapp.students
	WHERE user_id = $1
	`
	db := r.dbFromContext(ctx)
	row := db.QueryRow(ctx, query, userID)

	var m StudentModel
	err := row.Scan(
		&m.ID,
		&m.Version,
		&m.UserID,
		&m.GroupID,
		&m.FIO,
		&m.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Student{}, fmt.Errorf(
				"student by user_id not found: %w",
				core_errors.ErrNotFound,
			)
		}
		return domain.Student{}, fmt.Errorf("scan error: %w", err)
	}

	return domain.NewStudent(
		m.ID,
		m.Version,
		m.UserID,
		m.GroupID,
		m.FIO,
		m.PhoneNumber,
	), nil
}
