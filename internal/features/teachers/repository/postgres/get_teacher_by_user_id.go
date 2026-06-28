package teachers_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *TeachersRepository) GetTeacherByUserID(
	ctx context.Context,
	userID int,
) (domain.Teacher, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT
		id,
		version,
		user_id,
		fio,
		phone_number
	FROM myapp.teachers
	WHERE user_id = $1;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		userID,
	)

	var model TeacherModel

	err := row.Scan(
		&model.ID,
		&model.Version,
		&model.UserID,
		&model.FIO,
		&model.PhoneNumber,
	)
	if err != nil {
		return domain.Teacher{}, fmt.Errorf(
			"scan teacher: %w",
			err,
		)
	}

	return domain.NewTeacher(
		model.ID,
		model.Version,
		model.UserID,
		model.FIO,
		model.PhoneNumber,
	), nil
}
