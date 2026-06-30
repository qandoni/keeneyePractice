package admin_service

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_postgres "github.com/qandoni/keeneyePractice/internal/core/repository/postgres"
	users_contracts "github.com/qandoni/keeneyePractice/internal/features/users/contracts"
)

type AdminService struct {
	usersService    UsersService
	studentsService StudentsService
	teachersService TeachersService
	txManager       core_postgres.TransactionManager
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
}

func NewAdminService(
	usersService UsersService,
	studentsService StudentsService,
	teachersService TeachersService,
	txManager core_postgres.TransactionManager,
) *AdminService {
	return &AdminService{
		usersService:    usersService,
		studentsService: studentsService,
		teachersService: teachersService,
		txManager:       txManager,
	}
}
