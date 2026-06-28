package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	core_logger "github.com/qandoni/keeneyePractice/internal/core/logger"
	core_password "github.com/qandoni/keeneyePractice/internal/core/password"
	core_pgx_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/qandoni/keeneyePractice/internal/core/transport/http/middleware"
	core_http_server "github.com/qandoni/keeneyePractice/internal/core/transport/http/server"
	admin_service "github.com/qandoni/keeneyePractice/internal/features/admin/service"
	admin_transport_http "github.com/qandoni/keeneyePractice/internal/features/admin/transport/http"
	auth_jwt "github.com/qandoni/keeneyePractice/internal/features/auth/jwt"
	auth_service "github.com/qandoni/keeneyePractice/internal/features/auth/service"
	auth_http_transport "github.com/qandoni/keeneyePractice/internal/features/auth/transport/http"
	groups_postgres_repository "github.com/qandoni/keeneyePractice/internal/features/groups/repository/postgres"
	groups_service "github.com/qandoni/keeneyePractice/internal/features/groups/service"
	student_policy "github.com/qandoni/keeneyePractice/internal/features/students/policy"
	students_postgres_repository "github.com/qandoni/keeneyePractice/internal/features/students/repository/postgres"
	students_service "github.com/qandoni/keeneyePractice/internal/features/students/service"
	students_transport_http "github.com/qandoni/keeneyePractice/internal/features/students/transport/http"
	teachers_postgres_repository "github.com/qandoni/keeneyePractice/internal/features/teachers/repository/postgres"
	teachers_service "github.com/qandoni/keeneyePractice/internal/features/teachers/service"
	users_postgres_repository "github.com/qandoni/keeneyePractice/internal/features/users/repository/postgres"
	users_service "github.com/qandoni/keeneyePractice/internal/features/users/service"
	users_http_transport "github.com/qandoni/keeneyePractice/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init app logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing postgres connection pool")
	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "students"))
	studentsRepository := students_postgres_repository.NewStudentsRepository(pool)
	logger.Debug("initializing feature", zap.String("feature", "teachers"))
	teachersRepository := teachers_postgres_repository.NewUsersRepository(pool)
	teachersService := teachers_service.NewTeachersService(teachersRepository)
	studentsPolicy := student_policy.NewStudentAccessPolicy(studentsRepository, teachersRepository)
	studentsService := students_service.NewStudentsService(studentsRepository, studentsPolicy)
	studentsTransportHTTP := students_transport_http.NewStudentsHTTPHandler(studentsService)

	hasher := core_password.NewBcryptHasher()
	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository, hasher)
	usersTransportHTTP := users_http_transport.NewUsersHTTPHandler(usersService)

	logger.Debug("initializing feature", zap.String("feature", "groups"))
	groupsRepository := groups_postgres_repository.NewGroupsRepository(pool)
	groupsService := groups_service.NewGroupsService(groupsRepository)

	logger.Debug("initializing feature", zap.String("feature", "admin"))
	adminService := admin_service.NewAdminService(usersService, studentsService, teachersService, groupsService)
	adminTransportHTTP := admin_transport_http.NewAdminHTTPHandler(adminService)

	passwordHasher := core_password.NewBcryptHasher()
	logger.Debug("initializing feature", zap.String("feature", "auth"))
	jwtManager := auth_jwt.NewJWTManager("my-secret-key")
	authService := auth_service.NewAuthService(usersRepository, passwordHasher, jwtManager)
	authTransportHTTP := auth_http_transport.NewAuthHTTPHandler(authService)

	logger.Debug("initializing HTTP server")
	server := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
	)

	server.Engine().Use(
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		gin.Recovery(),
	)

	core_http_server.RegisterRoutes(
		server.Engine(),
		authTransportHTTP,
		adminTransportHTTP,
		studentsTransportHTTP,
		usersTransportHTTP,
		jwtManager,
	)

	if err := server.Run(ctx); err != nil {
		logger.Error(
			"http server stopped",
			zap.Error(err),
		)
	}
}
