package students_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *StudentsRepository) GetStudents(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Student, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, user_id, group_id, fio, phone_number
	FROM myapp.students
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
		return nil, fmt.Errorf("select students: %w", err)
	}
	defer rows.Close()

	var studentModels []StudentModel

	for rows.Next() {
		var studentModel StudentModel
		err := rows.Scan(
			&studentModel.ID,
			&studentModel.Version,
			&studentModel.UserID,
			&studentModel.GroupID,
			&studentModel.FIO,
			&studentModel.PhoneNumber,
		)
		if err != nil {
			return nil, fmt.Errorf("scan students: %w", err)
		}
		studentModels = append(studentModels, studentModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}
	studentDomains := studentDomainsFromModels(studentModels)
	return studentDomains, nil
}
