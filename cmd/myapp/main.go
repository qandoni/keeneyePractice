package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/qandoni/keeneyePractice/internal/core/logger"
	core_pgx_pool "github.com/qandoni/keeneyePractice/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/qandoni/keeneyePractice/internal/core/transport/http/middleware"
	core_http_server "github.com/qandoni/keeneyePractice/internal/core/transport/http/server"
	students_postgres_repository "github.com/qandoni/keeneyePractice/internal/features/students/repository/postgres"
	students_service "github.com/qandoni/keeneyePractice/internal/features/students/service"
	students_transport_http "github.com/qandoni/keeneyePractice/internal/features/students/transport/http"
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
	studentsService := students_service.NewStudentsService(studentsRepository)
	studentsTransportHTTP := students_transport_http.NewStudentsHTTPHandler(studentsService)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.RegisterRoutes(studentsTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouterV1)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
