package students_transport_http

import (
	"net/http"

	core_logger "github.com/qandoni/keeneyePractice/internal/core/logger"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

type GetStudentResponse StudentDTOResponse

func (h *StudentsHTTPHandler) GetStudent(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	studentID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get studentID path value",
		)
		return
	}
	student, err := h.studentsService.GetStudent(ctx, studentID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)
		return
	}
	response := GetStudentResponse(studentDTOFromDomain(student))
	responseHandler.JSONResponse(response, http.StatusOK)
}
