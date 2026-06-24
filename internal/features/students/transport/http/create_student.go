package students_transport_http

import (
	"net/http"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_logger "github.com/qandoni/keeneyePractice/internal/core/logger"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

type CreateStudentRequest struct {
	FIO          string `json:"fio" validate:"required,min=3,max=100"`
	StudentGroup string `json:"student_group" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required,min=10,max=15,startswith=+"`
}

type CreateStudentResponse StudentDTOResponse

func (h *StudentsHTTPHandler) CreateStudent(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateStudentRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	studentDomain := domainFromDTO(request)
	studentDomain, err := h.studentsService.CreateStudent(ctx, studentDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := CreateStudentResponse(studentDTOFromDomain(studentDomain))
	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateStudentRequest) domain.Student {
	return domain.NewStudentUninitialized(dto.FIO, dto.StudentGroup, dto.PhoneNumber)
}
