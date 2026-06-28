package core_http_server

import (
	"github.com/gin-gonic/gin"

	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_http_middleware "github.com/qandoni/keeneyePractice/internal/core/transport/http/middleware"
	admin_transport_http "github.com/qandoni/keeneyePractice/internal/features/admin/transport/http"
	auth_transport_http "github.com/qandoni/keeneyePractice/internal/features/auth/transport/http"
	students_transport_http "github.com/qandoni/keeneyePractice/internal/features/students/transport/http"
	users_transport_http "github.com/qandoni/keeneyePractice/internal/features/users/transport/http"
)

func RegisterRoutes(
	engine *gin.Engine,

	authHandler *auth_transport_http.AuthHTTPHandler,
	adminHandler *admin_transport_http.AdminHTTPHandler,
	studentsHandler *students_transport_http.StudentsHTTPHandler,
	usersHandler *users_transport_http.UsersHTTPHandler,

	parser core_http_middleware.TokenParser,
) {
	jwt := core_http_middleware.JWT(parser)

	api := engine.Group("/api/v1")

	auth := api.Group("/auth")
	authHandler.Register(auth)

	admin := api.Group("/admin")
	admin.Use(
		jwt,
		core_http_middleware.Role(enum.RoleAdmin),
	)
	adminHandler.Register(admin)

	students := api.Group("/students")
	students.Use(jwt)
	studentsHandler.Register(students)

	users := api.Group("/users")
	users.Use(
		jwt,
		core_http_middleware.Role(enum.RoleAdmin),
	)
	usersHandler.Register(users)
}
