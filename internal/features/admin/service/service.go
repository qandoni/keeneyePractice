package admin_service

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	group_contracts "github.com/qandoni/keeneyePractice/internal/features/groups/contracts"
	users_contracts "github.com/qandoni/keeneyePractice/internal/features/users/contracts"
)

type AdminService struct {
	usersService    UsersService
	studentsService StudentsService
	teachersService TeachersService
	groupsService   GroupsService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		input users_contracts.CreateUserInput,
	) (domain.User, error)
}

type StudentsService interface {
	CreateStudent(
		ctx context.Context,
		student domain.Student,
	) (domain.Student, error)
}

type TeachersService interface {
	CreateTeacher(
		ctx context.Context,
		teacher domain.Teacher,
	) (domain.Teacher, error)
	AssignToGroup(
		ctx context.Context,
		teacherID int,
		groupID int,
	) error
	RemoveFromGroup(
		ctx context.Context,
		teacherID int,
		groupID int,
	) error
	GetTeacherByUserID(
		ctx context.Context,
		userID int,
	) (domain.Teacher, error)
	GetTeachers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.Teacher, error)
}

type GroupsService interface {
	CreateGroup(
		ctx context.Context,
		input group_contracts.CreateGroupInput,
	) (domain.Group, error)
}

func NewAdminService(
	usersService UsersService,
	studentsService StudentsService,
	teachersService TeachersService,
	groupsService GroupsService,
) *AdminService {
	return &AdminService{
		usersService:    usersService,
		studentsService: studentsService,
		teachersService: teachersService,
		groupsService:   groupsService,
	}
}
