package students_postgres_repository

import (
	"context"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

func (r *StudentsRepository) CreateStudent(
	ctx context.Context,
	student domain.Student,
) (domain.Student, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	query := `
	INSERT INTO myapp.students (fio, student_group, phone_number)
	VALUES ($1, $2, $3)
	RETURNING id, version, fio, student_group, phone_number;
	`

	row := r.pool.QueryRow(ctx, query, student.FIO, student.StudentGroup, student.PhoneNumber)

	var studentModel StudentModel

	err := row.Scan(
		&studentModel.ID,
		&studentModel.Version,
		&studentModel.FIO,
		&studentModel.StudentGroup,
		&studentModel.PhoneNumber,
	)
	if err != nil {
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
