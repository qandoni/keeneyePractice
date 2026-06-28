package teachers_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *TeachersRepository) CreateTeacher(
	ctx context.Context,
	teacher domain.Teacher,
) (domain.Teacher, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO myapp.teachers(user_id, fio, phone_number)
	VALUES ($1, $2, $3)
	RETURNING id, version, user_id, fio, phone_number;
	`

	row := r.pool.QueryRow(
		ctx, query,
		teacher.UserID,
		teacher.FIO,
		teacher.PhoneNumber,
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
