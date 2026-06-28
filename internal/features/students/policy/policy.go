package student_policy

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

type StudentAccessPolicy struct {
	studentsRepository StudentsRepository
	teachersRepository TeachersRepository
}

type StudentsRepository interface {
	GetStudent(
		ctx context.Context,
		id int,
	) (domain.Student, error)

	GetStudentByUserID(
		ctx context.Context,
		userID int,
	) (domain.Student, error)
}

type TeachersRepository interface {
	GetTeacherByUserID(
		ctx context.Context,
		userID int,
	) (domain.Teacher, error)

	TeacherHasGroup(
		ctx context.Context,
		teacherID int,
		groupID int,
	) (bool, error)

	GetTeacherGroupIDs(
		ctx context.Context,
		teacherID int,
	) ([]int, error)
}

func NewStudentAccessPolicy(
	studentsRepository StudentsRepository,
	teachersRepository TeachersRepository,
) *StudentAccessPolicy {
	return &StudentAccessPolicy{
		studentsRepository: studentsRepository,
		teachersRepository: teachersRepository,
	}
}
