package core_http_server

import (
	"github.com/gin-gonic/gin"

	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_http_middleware "github.com/qandoni/keeneyePractice/internal/core/transport/http/middleware"
	admin_transport_http "github.com/qandoni/keeneyePractice/internal/features/admin/transport/http"
	auth_transport_http "github.com/qandoni/keeneyePractice/internal/features/auth/transport/http"
	groups_http_transport "github.com/qandoni/keeneyePractice/internal/features/groups/transport/http"
	students_transport_http "github.com/qandoni/keeneyePractice/internal/features/students/transport/http"
	teachers_transport_http "github.com/qandoni/keeneyePractice/internal/features/teachers/transport/http"
	users_transport_http "github.com/qandoni/keeneyePractice/internal/features/users/transport/http"
)

func RegisterRoutes(
	engine *gin.Engine,

	authHandler *auth_transport_http.AuthHTTPHandler,
	adminHandler *admin_transport_http.AdminHTTPHandler,
	studentsHandler *students_transport_http.StudentsHTTPHandler,
	teachersHandler *teachers_transport_http.TeachersHTTPHandler,
	groupsHandler *groups_http_transport.GroupsHTTPHandler,
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

	teachers := api.Group("/teachers")
	teachers.Use(
		jwt,
		core_http_middleware.Role(enum.RoleAdmin),
	)
	teachersHandler.Register(teachers)

	groups := api.Group("/groups")
	groups.Use(
		jwt,
		core_http_middleware.Role(enum.RoleAdmin),
	)
	groupsHandler.Register(groups)

	users := api.Group("/users")
	users.Use(
		jwt,
		core_http_middleware.Role(enum.RoleAdmin),
	)
	usersHandler.Register(users)
}
