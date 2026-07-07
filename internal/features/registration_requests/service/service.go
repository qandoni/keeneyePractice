package registration_service

import (
	"context"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_postgres "github.com/qandoni/keeneyePractice/internal/core/repository/postgres"
	users_contracts "github.com/qandoni/keeneyePractice/internal/features/users/contracts"
)

type RegistrationRequestService struct {
	registrationRepository RegistrationRequestRepository
	groupsService          GroupsService
	usersService           UsersService
	studentsService        StudentsService
	teachersService        TeachersService
	tokenGenerator         TokenGenerator
	sha256Hasher           Sha256Hasher
	emailSender            EmailSender
	txManager              core_postgres.TransactionManager
}

type RegistrationRequestRepository interface {
	Create(
		ctx context.Context,
		requets domain.RegistrationRequest,
	) error
	UpdateStatus(
		ctx context.Context,
		id int,
		status enum.RegistrationStatus,
	) error
	GetByTokenHash(
		ctx context.Context,
		hash string,
	) (domain.RegistrationRequest, error)
	ExpireRequests(
		ctx context.Context,
	) error
}

type PasswordHasher interface {
	Hash(
		password string,
	) (string, error)
	Compare(
		hash string,
		password string,
	) error
}

type EmailSender interface {
	SendRegistrationEmail(ctx context.Context, to string, subject string, body string) error
}

type GroupsService interface {
	GetGroupByName(
		ctx context.Context,
		name string,
	) (domain.Group, error)
}
type UsersService interface {
	GetUserByEmail(
		ctx context.Context,
		email string,
	) (domain.User, error)
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

type TokenGenerator interface {
	Generate() (string, error)
}
type Sha256Hasher interface {
	Hash(value string) string
}

func NewRegistrationRequestsService(
	registrationRepository RegistrationRequestRepository,
	groupsService GroupsService,
	usersService UsersService,
	studentsService StudentsService,
	teacehrsService TeachersService,
	tokenGenerator TokenGenerator,
	sha256Hasher Sha256Hasher,
	emailSender EmailSender,
	txManager core_postgres.TransactionManager,
) *RegistrationRequestService {
	return &RegistrationRequestService{
		registrationRepository: registrationRepository,
		groupsService:          groupsService,
		usersService:           usersService,
		studentsService:        studentsService,
		teachersService:        teacehrsService,
		tokenGenerator:         tokenGenerator,
		sha256Hasher:           sha256Hasher,
		emailSender:            emailSender,
		txManager:              txManager,
	}
}
