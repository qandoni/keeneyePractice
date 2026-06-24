package students_transport_http

import (
	"context"
	"net/http"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_http_server "github.com/qandoni/keeneyePractice/internal/core/transport/http/server"
)

type StudentsHTTPHandler struct {
	studentsService StudentsService
}

type StudentsService interface {
	CreateStudent(
		ctx context.Context,
		student domain.Student,
	) (domain.Student, error)
	GetStudent(
		ctx context.Context,
		id int,
	) (domain.Student, error)
	PatchStudent(
		ctx context.Context,
		id int,
		patch domain.StudentPatch,
	) (domain.Student, error)
	DeleteStudent(
		ctx context.Context,
		id int,
	) error
}

func NewStudentsHTTPHandler(
	studentsService StudentsService,
) *StudentsHTTPHandler {
	return &StudentsHTTPHandler{
		studentsService: studentsService,
	}
}

func (h *StudentsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/students",
			Handler: h.CreateStudent,
		},
		{
			Method:  http.MethodGet,
			Path:    "/students/{id}",
			Handler: h.GetStudent,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/students/{id}",
			Handler: h.PatchStudent,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/students/{id}",
			Handler: h.DeleteStudent,
		},
	}
}
