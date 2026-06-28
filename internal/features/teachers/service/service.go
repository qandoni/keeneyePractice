package teachers_service

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

type TeachersService struct {
	teachersRepository TeachersRepository
}

type TeachersRepository interface {
	CreateTeacher(
		ctx context.Context,
		teacher domain.Teacher,
	) (domain.Teacher, error)
	GetTeacher(
		ctx context.Context,
		id int,
	) (domain.Teacher, error)
	GetTeachers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.Teacher, error)
	PatchTeacher(
		ctx context.Context,
		id int,
		teacher domain.Teacher,
	) (domain.Teacher, error)
	DeleteTeacher(
		ctx context.Context,
		id int,
	) error
	GetTeacherByUserID(
		ctx context.Context,
		userID int,
	) (domain.Teacher, error)
	AssignToGroup(
		ctx context.Context,
		assignment domain.TeacherGroupAssignment,
	) error
	RemoveFromGroup(
		ctx context.Context,
		teacherID int,
		groupID int,
	) error
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

func NewTeachersService(
	teachersRepository TeachersRepository,
) *TeachersService {
	return &TeachersService{
		teachersRepository: teachersRepository,
	}
}
