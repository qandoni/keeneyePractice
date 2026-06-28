package teachers_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	core_postgres_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool"
)

func (r *TeachersRepository) GetTeacher(
	ctx context.Context,
	id int,
) (domain.Teacher, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, user_id, fio, phone_number
	FROM myapp.teachers
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)
	var TeacherModel TeacherModel

	err := row.Scan(
		&TeacherModel.ID,
		&TeacherModel.Version,
		&TeacherModel.UserID,
		&TeacherModel.FIO,
		&TeacherModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Teacher{}, fmt.Errorf("teacher with id='%d': %w", id, core_errors.ErrNotFound)
		}
		return domain.Teacher{}, fmt.Errorf("scan error: %w", err)
	}

	teacherDomain := domain.NewTeacher(
		TeacherModel.ID,
		TeacherModel.Version,
		TeacherModel.UserID,
		TeacherModel.FIO,
		TeacherModel.PhoneNumber,
	)
	return teacherDomain, nil
}
