package teachers_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *TeachersRepository) GetTeachers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Teacher, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, user_id, fio, phone_number
	FROM myapp.teachers
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("select teachers: %w", err)
	}
	defer rows.Close()

	var teacherModels []TeacherModel

	for rows.Next() {
		var teacherModel TeacherModel
		err := rows.Scan(
			&teacherModel.ID,
			&teacherModel.Version,
			&teacherModel.UserID,
			&teacherModel.FIO,
			&teacherModel.PhoneNumber,
		)
		if err != nil {
			return nil, fmt.Errorf("scan teacher: %w", err)
		}
		teacherModels = append(teacherModels, teacherModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}
	teacherDomains := teachersDomainsFromModels(teacherModels)
	return teacherDomains, nil
}
