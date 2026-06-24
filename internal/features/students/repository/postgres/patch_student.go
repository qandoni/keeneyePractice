package students_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	core_postgres_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool"
)

func (r *StudentsRepository) PatchStudent(
	ctx context.Context,
	id int,
	student domain.Student,
) (domain.Student, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE myapp.students
	SET
		fio=$1,
		student_group=$2,
		phone_number=$3,
		version=version+1
	WHERE id=$4 AND version=$5
	RETURNING
		id,
		version,
		fio,
		student_group,
		phone_number;
	`
	row := r.pool.QueryRow(
		ctx,
		query,
		student.FIO,
		student.StudentGroup,
		student.PhoneNumber,
		id,
		student.Version,
	)

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
			return domain.Student{}, fmt.Errorf(
				"student with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}
		return domain.Student{}, fmt.Errorf("scan error: %w", err)
	}
	studentDomain := domain.NewStudent(
		student.ID,
		student.Version,
		student.FIO,
		student.StudentGroup,
		student.PhoneNumber,
	)
	return studentDomain, nil
}
