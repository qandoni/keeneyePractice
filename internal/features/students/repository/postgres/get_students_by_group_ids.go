package students_postgres_repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *StudentsRepository) GetStudentsByGroupIDs(
	ctx context.Context,
	groupIDs []int,
	limit *int,
	offset *int,
) ([]domain.Student, error) {

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
		SELECT id, version, user_id, group_id, fio, phone_number
		FROM myapp.students
		WHERE group_id = ANY($1)
		ORDER BY id
	`

	args := []any{groupIDs}

	if limit != nil {
		query += " LIMIT $" + strconv.Itoa(len(args)+1)
		args = append(args, *limit)
	}

	if offset != nil {
		query += " OFFSET $" + strconv.Itoa(len(args)+1)
		args = append(args, *offset)
	}

	db := r.dbFromContext(ctx)
	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query students by group ids: %w", err)
	}
	defer rows.Close()

	var result []domain.Student

	for rows.Next() {
		var s domain.Student

		err := rows.Scan(
			&s.ID,
			&s.Version,
			&s.UserID,
			&s.GroupID,
			&s.FIO,
			&s.PhoneNumber,
		)
		if err != nil {
			return nil, fmt.Errorf("scan student: %w", err)
		}

		result = append(result, s)
	}

	return result, nil
}
