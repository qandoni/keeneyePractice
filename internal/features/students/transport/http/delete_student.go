package students_transport_http

import (
	"net/http"

	core_logger "github.com/qandoni/keeneyePractice/internal/core/logger"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

func (h *StudentsHTTPHandler) DeleteStudent(rw http.ResponseWriter, r *http.Request) {
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
	if err := h.studentsService.DeleteStudent(ctx, studentID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete student",
		)
		return
	}
	responseHandler.NoContentResponse()
}
