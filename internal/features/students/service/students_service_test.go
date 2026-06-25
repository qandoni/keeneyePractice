package students_service

import (
	"context"
	"errors"
	"testing"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	"github.com/stretchr/testify/require"
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

func TestCreateStudent_Success(t *testing.T) {
	repo := &MockStudentsRepository{
		CreateStudentFunc: func(
			ctx context.Context,
			student domain.Student,
		) (domain.Student, error) {

			return domain.NewStudent(
				1,
				1,
				student.FIO,
				student.StudentGroup,
				student.PhoneNumber,
			), nil
		},
	}

	service := NewStudentsService(repo)

	student := domain.NewStudentUninitialized(
		"Ivan Ivanov",
		"PI-21",
		"+79999999999",
	)

	result, err := service.CreateStudent(
		context.Background(),
		student,
	)

	require.NoError(t, err)
	require.Equal(t, 1, result.ID)
}

func TestCreateStudent_InvalidStudent(t *testing.T) {
	repo := &MockStudentsRepository{
		CreateStudentFunc: func(
			ctx context.Context,
			student domain.Student,
		) (domain.Student, error) {

			t.Fatal("repository should not be called")

			return domain.Student{}, nil
		},
	}

	service := NewStudentsService(repo)

	student := domain.NewStudentUninitialized(
		"ab",
		"PI-21",
		"+79999999999",
	)

	_, err := service.CreateStudent(
		context.Background(),
		student,
	)

	require.Error(t, err)
}

func TestCreateStudent_RepositoryError(t *testing.T) {
	repo := &MockStudentsRepository{
		CreateStudentFunc: func(
			ctx context.Context,
			student domain.Student,
		) (domain.Student, error) {

			return domain.Student{}, errors.New("db error")
		},
	}

	service := NewStudentsService(repo)

	student := domain.NewStudentUninitialized(
		"Ivan Ivanov",
		"PI-21",
		"+79999999999",
	)

	_, err := service.CreateStudent(
		context.Background(),
		student,
	)

	require.Error(t, err)
}

func TestGetStudent_Success(t *testing.T) {
	repo := &MockStudentsRepository{
		GetStudentFunc: func(
			ctx context.Context,
			id int,
		) (domain.Student, error) {

			return domain.NewStudent(
				id,
				1,
				"Ivan",
				"PI-21",
				"+79999999999",
			), nil
		},
	}

	service := NewStudentsService(repo)

	student, err := service.GetStudent(
		context.Background(),
		1,
	)

	require.NoError(t, err)
	require.Equal(t, 1, student.ID)
}

func TestGetStudent_RepositoryError(t *testing.T) {
	repo := &MockStudentsRepository{
		GetStudentFunc: func(
			ctx context.Context,
			id int,
		) (domain.Student, error) {

			return domain.Student{}, errors.New("db error")
		},
	}

	service := NewStudentsService(repo)

	_, err := service.GetStudent(
		context.Background(),
		1,
	)

	require.Error(t, err)
}

func TestGetStudents_NegativeOffset(t *testing.T) {
	repo := &MockStudentsRepository{
		GetStudentsFunc: func(
			ctx context.Context,
			limit *int,
			offset *int,
		) ([]domain.Student, error) {

			t.Fatal("repository should not be called")

			return nil, nil
		},
	}

	service := NewStudentsService(repo)

	offset := -1

	_, err := service.GetStudents(
		context.Background(),
		nil,
		&offset,
	)

	require.Error(t, err)
	require.ErrorIs(t, err, core_errors.ErrInvalidArgument)
}

func TestDeleteStudent_Success(t *testing.T) {
	repo := &MockStudentsRepository{
		DeleteStudentFunc: func(
			ctx context.Context,
			id int,
		) error {

			require.Equal(t, 1, id)

			return nil
		},
	}

	service := NewStudentsService(repo)

	err := service.DeleteStudent(
		context.Background(),
		1,
	)

	require.NoError(t, err)
}

func TestDeleteStudent_RepositoryError(t *testing.T) {
	repo := &MockStudentsRepository{
		DeleteStudentFunc: func(
			ctx context.Context,
			id int,
		) error {

			return errors.New("db error")
		},
	}

	service := NewStudentsService(repo)

	err := service.DeleteStudent(
		context.Background(),
		1,
	)

	require.Error(t, err)
}

func TestPatchStudent_Success(t *testing.T) {
	repo := &MockStudentsRepository{
		GetStudentFunc: func(
			ctx context.Context,
			id int,
		) (domain.Student, error) {

			return domain.NewStudent(
				1,
				1,
				"Old Name",
				"PI-21",
				"+79999999999",
			), nil
		},

		PatchStudentFunc: func(
			ctx context.Context,
			id int,
			student domain.Student,
		) (domain.Student, error) {

			return student, nil
		},
	}

	service := NewStudentsService(repo)

	newName := "New Name"

	patch := domain.NewStudentPatch(
		domain.Nullable[string]{
			Set:   true,
			Value: &newName,
		},
		domain.Nullable[string]{},
		domain.Nullable[string]{},
	)

	result, err := service.PatchStudent(
		context.Background(),
		1,
		patch,
	)

	require.NoError(t, err)
	require.Equal(t, "New Name", result.FIO)
}
func TestPatchStudent_GetStudentError(t *testing.T) {
	repo := &MockStudentsRepository{
		GetStudentFunc: func(
			ctx context.Context,
			id int,
		) (domain.Student, error) {

			return domain.Student{}, errors.New("db error")
		},
	}

	service := NewStudentsService(repo)

	_, err := service.PatchStudent(
		context.Background(),
		1,
		domain.StudentPatch{},
	)

	require.Error(t, err)
}

func TestPatchStudent_InvalidPatch(t *testing.T) {
	repo := &MockStudentsRepository{
		GetStudentFunc: func(
			ctx context.Context,
			id int,
		) (domain.Student, error) {

			return domain.NewStudent(
				1,
				1,
				"Ivan",
				"PI-21",
				"+79999999999",
			), nil
		},
	}

	service := NewStudentsService(repo)

	patch := domain.NewStudentPatch(
		domain.Nullable[string]{
			Set:   true,
			Value: nil,
		},
		domain.Nullable[string]{},
		domain.Nullable[string]{},
	)

	_, err := service.PatchStudent(
		context.Background(),
		1,
		patch,
	)

	require.Error(t, err)
}

func TestPatchStudent_PatchRepositoryError(t *testing.T) {
	repo := &MockStudentsRepository{
		GetStudentFunc: func(
			ctx context.Context,
			id int,
		) (domain.Student, error) {

			return domain.NewStudent(
				1,
				1,
				"Ivan",
				"PI-21",
				"+79999999999",
			), nil
		},

		PatchStudentFunc: func(
			ctx context.Context,
			id int,
			student domain.Student,
		) (domain.Student, error) {

			return domain.Student{}, errors.New("db error")
		},
	}

	service := NewStudentsService(repo)

	newName := "New Name"

	patch := domain.NewStudentPatch(
		domain.Nullable[string]{
			Set:   true,
			Value: &newName,
		},
		domain.Nullable[string]{},
		domain.Nullable[string]{},
	)

	_, err := service.PatchStudent(
		context.Background(),
		1,
		patch,
	)

	require.Error(t, err)
}

func TestPatchStudent_ApplyPatch(t *testing.T) {
	var receivedStudent domain.Student

	repo := &MockStudentsRepository{
		GetStudentFunc: func(
			ctx context.Context,
			id int,
		) (domain.Student, error) {

			return domain.NewStudent(
				1,
				1,
				"Old Name",
				"PI-21",
				"+79999999999",
			), nil
		},

		PatchStudentFunc: func(
			ctx context.Context,
			id int,
			student domain.Student,
		) (domain.Student, error) {

			receivedStudent = student

			return student, nil
		},
	}

	service := NewStudentsService(repo)

	newName := "New Name"

	patch := domain.NewStudentPatch(
		domain.Nullable[string]{
			Set:   true,
			Value: &newName,
		},
		domain.Nullable[string]{},
		domain.Nullable[string]{},
	)

	_, err := service.PatchStudent(
		context.Background(),
		1,
		patch,
	)

	require.NoError(t, err)

	require.Equal(
		t,
		"New Name",
		receivedStudent.FIO,
	)
}
