package students_service

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

type MockStudentsRepository struct {
	CreateStudentFunc func(
		ctx context.Context,
		student domain.Student,
	) (domain.Student, error)

	GetStudentFunc func(
		ctx context.Context,
		id int,
	) (domain.Student, error)

	PatchStudentFunc func(
		ctx context.Context,
		id int,
		student domain.Student,
	) (domain.Student, error)

	DeleteStudentFunc func(
		ctx context.Context,
		id int,
	) error

	GetStudentsFunc func(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.Student, error)
}

func (m *MockStudentsRepository) CreateStudent(
	ctx context.Context,
	student domain.Student,
) (domain.Student, error) {
	return m.CreateStudentFunc(ctx, student)
}

func (m *MockStudentsRepository) GetStudent(
	ctx context.Context,
	id int,
) (domain.Student, error) {
	return m.GetStudentFunc(ctx, id)
}

func (m *MockStudentsRepository) PatchStudent(
	ctx context.Context,
	id int,
	student domain.Student,
) (domain.Student, error) {
	return m.PatchStudentFunc(ctx, id, student)
}

func (m *MockStudentsRepository) DeleteStudent(
	ctx context.Context,
	id int,
) error {
	return m.DeleteStudentFunc(ctx, id)
}

func (m *MockStudentsRepository) GetStudents(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Student, error) {
	return m.GetStudentsFunc(ctx, limit, offset)
}
