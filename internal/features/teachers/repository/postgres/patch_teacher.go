package teachers_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	core_postgres_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool"
)

func (r *TeachersRepository) PatchTeacher(
	ctx context.Context,
	id int,
	teacher domain.Teacher,
) (domain.Teacher, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	UPDATE myapp.teachers
	SET
		fio=$1,
		phone_number=$2,
		version=version+1
	WHERE id=$3 AND version=$4
	RETURNING
		id,
		version,
		user_id,
		fio,
		phone_number;
	`
	db := r.dbFromContext(ctx)
	row := db.QueryRow(
		ctx,
		query,
		teacher.FIO,
		teacher.PhoneNumber,
		id,
		teacher.Version,
	)

	var teacherModel TeacherModel
	err := row.Scan(
		&teacherModel.ID,
		&teacherModel.Version,
		&teacherModel.UserID,
		&teacherModel.FIO,
		&teacherModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Teacher{}, fmt.Errorf(
				"teacher with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}
		return domain.Teacher{}, fmt.Errorf("scan error: %w", err)
	}
	teacherDomain := domain.NewTeacher(
		teacherModel.ID,
		teacherModel.Version,
		teacherModel.UserID,
		teacherModel.FIO,
		teacherModel.PhoneNumber,
	)
	return teacherDomain, nil
}
