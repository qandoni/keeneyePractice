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
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `
	UPDATE myapp.students
	SET
		fio=$1,
		phone_number=$2,
		version=version+1
	WHERE id=$3 AND version=$4
	RETURNING
		id,
		version,
		user_id,
		group_id,
		fio,
		phone_number;
	`
	db := r.dbFromContext(ctx)
	row := db.QueryRow(
		ctx,
		query,
		student.FIO,
		student.PhoneNumber,
		id,
		student.Version,
	)

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
		studentModel.ID,
		studentModel.Version,
		studentModel.UserID,
		studentModel.GroupID,
		studentModel.FIO,
		studentModel.PhoneNumber,
	)
	return studentDomain, nil
}
