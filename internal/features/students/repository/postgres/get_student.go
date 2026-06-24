package students_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	core_postgres_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool"
)

func (r *StudentsRepository) GetStudent(
	ctx context.Context,
	id int,
) (domain.Student, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, fio, student_group, phone_number
	FROM myapp.students
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)
	var studentModel StudentModel

	err := row.Scan(
		&studentModel.ID,
		&studentModel.Version,
		&studentModel.FIO,
		&studentModel.StudentGroup,
		&studentModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Student{}, fmt.Errorf("user with id='%d': %w", id, core_errors.ErrNotFound)
		}
		return domain.Student{}, fmt.Errorf("scan error: %w", err)
	}

	studentDomain := domain.NewStudent(
		studentModel.ID,
		studentModel.Version,
		studentModel.FIO,
		studentModel.StudentGroup,
		studentModel.PhoneNumber,
	)
	return studentDomain, nil
}
