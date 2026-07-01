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
	core_password_hash "github.com/qandoni/keeneyePractice/internal/core/password/hash"
	core_pgx_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/qandoni/keeneyePractice/internal/core/transport/http/middleware"
	core_http_server "github.com/qandoni/keeneyePractice/internal/core/transport/http/server"
	admin_service "github.com/qandoni/keeneyePractice/internal/features/admin/service"
	admin_transport_http "github.com/qandoni/keeneyePractice/internal/features/admin/transport/http"
	auth_jwt "github.com/qandoni/keeneyePractice/internal/features/auth/jwt"
	auth_refresh "github.com/qandoni/keeneyePractice/internal/features/auth/refresh"
	auth_postgres_repository "github.com/qandoni/keeneyePractice/internal/features/auth/repository/postgres"
	auth_service "github.com/qandoni/keeneyePractice/internal/features/auth/service"
	auth_http_transport "github.com/qandoni/keeneyePractice/internal/features/auth/transport/http"
	groups_postgres_repository "github.com/qandoni/keeneyePractice/internal/features/groups/repository/postgres"
	groups_service "github.com/qandoni/keeneyePractice/internal/features/groups/service"
	groups_http_transport "github.com/qandoni/keeneyePractice/internal/features/groups/transport/http"
	student_policy "github.com/qandoni/keeneyePractice/internal/features/students/policy"
	students_postgres_repository "github.com/qandoni/keeneyePractice/internal/features/students/repository/postgres"
	students_service "github.com/qandoni/keeneyePractice/internal/features/students/service"
	students_transport_http "github.com/qandoni/keeneyePractice/internal/features/students/transport/http"
	teachers_postgres_repository "github.com/qandoni/keeneyePractice/internal/features/teachers/repository/postgres"
	teachers_service "github.com/qandoni/keeneyePractice/internal/features/teachers/service"
	teachers_transport_http "github.com/qandoni/keeneyePractice/internal/features/teachers/transport/http"
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

	logger.Debug("initializing transaction manager")
	txManager := core_pgx_pool.NewTransactionManager(pool)

	logger.Debug("initializing feature", zap.String("feature", "students"))
	logger.Debug("initializing feature", zap.String("feature", "teachers"))
	studentsRepository := students_postgres_repository.NewStudentsRepository(pool, pool.OpTimeout())
	teachersRepository := teachers_postgres_repository.NewUsersRepository(pool, pool.OpTimeout())
	teachersService := teachers_service.NewTeachersService(teachersRepository)
	teachersTransportHTTP := teachers_transport_http.NewTeachersHTTPHandler(teachersService)
	studentsPolicy := student_policy.NewStudentAccessPolicy(studentsRepository, teachersRepository)
	studentsService := students_service.NewStudentsService(studentsRepository, studentsPolicy)
	studentsTransportHTTP := students_transport_http.NewStudentsHTTPHandler(studentsService)

	hasher := core_password.NewBcryptHasher()
	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool, pool.OpTimeout())
	usersService := users_service.NewUsersService(usersRepository, hasher)
	usersTransportHTTP := users_http_transport.NewUsersHTTPHandler(usersService)

	logger.Debug("initializing feature", zap.String("feature", "groups"))
	groupsRepository := groups_postgres_repository.NewGroupsRepository(pool, pool.OpTimeout())
	groupsService := groups_service.NewGroupsService(groupsRepository)
	groupsTransportHTTP := groups_http_transport.NewGroupsHTTPHandler(groupsService)

	logger.Debug("initializing feature", zap.String("feature", "admin"))
	adminService := admin_service.NewAdminService(usersService, studentsService, teachersService, txManager)
	adminTransportHTTP := admin_transport_http.NewAdminHTTPHandler(adminService)

	passwordHasher := core_password.NewBcryptHasher()
	logger.Debug("initializing feature", zap.String("feature", "auth"))
	refreshRepository := auth_postgres_repository.NewRefreshTokensRepository(pool, pool.OpTimeout())
	refreshGenerator := auth_refresh.NewGenerator()
	sha256Hasher := core_password_hash.NewSHA256Hasher()
	jwtManager := auth_jwt.NewJWTManager("my-secret-key")
	authService := auth_service.NewAuthService(usersRepository, refreshRepository, passwordHasher, sha256Hasher, jwtManager, refreshGenerator, txManager)
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
		teachersTransportHTTP,
		groupsTransportHTTP,
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
