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
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	query := `
	INSERT INTO myapp.students (user_id, group_id, fio, phone_number)
	VALUES ($1, $2, $3, $4)
	RETURNING id, version, user_id, group_id, fio, phone_number;
	`
	db := r.dbFromContext(ctx)
	row := db.QueryRow(ctx, query,
		student.UserID,
		student.GroupID,
		student.FIO,
		student.PhoneNumber)

	var studentModel StudentModel

	err := row.Scan(
		&studentModel.ID,
		&studentModel.Version,
		&studentModel.UserID,
		&studentModel.GroupID,
		&studentModel.FIO,
		&studentModel.PhoneNumber,
	)
	if err != nil {
		return domain.Student{}, fmt.Errorf("scan error: %w", err)
	}

	studentDomain := domain.NewStudent(
		studentModel.ID,
		studentModel.Version,
		studentModel.UserID,
		studentModel.GroupID,
		studentModel.FIO,
		studentModel.PhoneNumber,
	)
	return studentDomain, nil
}
