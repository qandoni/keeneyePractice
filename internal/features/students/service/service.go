package students_service

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	student_policy "github.com/qandoni/keeneyePractice/internal/features/students/policy"
)

type StudentsService struct {
	studentsRepository StudentsRepository
	policy             student_policy.StudentAccessPolicy
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
	GetStudentsByGroupIDs(
		ctx context.Context,
		groupIDs []int,
		limit *int,
		offset *int,
	) ([]domain.Student, error)
	GetStudentByUserID(
		ctx context.Context,
		userID int,
	) (domain.Student, error)
}

func NewStudentsService(
	studentsRepository StudentsRepository,
	studentPolicy *student_policy.StudentAccessPolicy,
) *StudentsService {
	return &StudentsService{
		studentsRepository: studentsRepository,
		policy:             *studentPolicy,
	}
}
