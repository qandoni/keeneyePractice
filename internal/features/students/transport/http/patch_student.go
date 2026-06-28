package students_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
	core_http_types "github.com/qandoni/keeneyePractice/internal/core/transport/http/types"
)

type PatchStudentRequest struct {
	FIO          core_http_types.Nullable[string] `json:"fio"`
	StudentGroup core_http_types.Nullable[string] `json:"student_group"`
	PhoneNumber  core_http_types.Nullable[string] `json:"phone_number"`
}

func (r *PatchStudentRequest) Validate() error {
	if r.FIO.Set {
		if r.FIO.Value == nil {
			return fmt.Errorf("`FIO` can't be NULL")
		}
		fioLen := len([]rune(*r.FIO.Value))
		if fioLen < 3 || fioLen > 100 {
			return fmt.Errorf("`FIO` must be between 3 and 100 symbols")
		}
	}
	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("`PhoneNumber` must be between 10 and 15 symbols")
			}
			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("`PhoneNumber` must start with '+' symbol ")
			}
		}
	}
	return nil
}

type PatchStudentResponse StudentDTOResponse

func (h *StudentsHTTPHandler) PatchStudent(c *gin.Context) {
	ctx := c.Request.Context()

	studentID, err := core_http_request.GetIntPathValue(c, "id")
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get int path value",
		)
		return
	}
	var request PatchStudentRequest

	if err := core_http_request.DecodeAndValidateRequest(c, &request); err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	studentPatch := studentPatchFromRequest(request)

	studentDomain, err := h.studentsService.PatchStudent(ctx, studentID, studentPatch)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to patch student",
		)
		return
	}
	response := PatchStudentResponse(studentDTOFromDomain(studentDomain))

	c.JSON(http.StatusOK, response)
}

func studentPatchFromRequest(request PatchStudentRequest) domain.StudentPatch {
	return domain.NewStudentPatch(
		request.FIO.ToDomain(),
		request.PhoneNumber.ToDomain(),
	)
}
