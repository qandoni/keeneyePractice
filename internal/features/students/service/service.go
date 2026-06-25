package students_service

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

type StudentsService struct {
	studentsRepository StudentsRepository
}

type StudentsRepository interface {
	CreateStudent(
		ctx context.Context,
		student domain.Student,
	) (domain.Student, error)
	GetStudent(
		ctx context.Context,
		id int,
	) (domain.Student, error)
	PatchStudent(
		ctx context.Context,
		id int,
		student domain.Student,
	) (domain.Student, error)
	DeleteStudent(
		ctx context.Context,
		id int,
	) error
	GetStudents(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.Student, error)
}

func NewStudentsService(
	studentsRepository StudentsRepository,
) *StudentsService {
	return &StudentsService{
		studentsRepository: studentsRepository,
	}
}
