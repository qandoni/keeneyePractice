package registration_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
)

type GetAllRequestsResponse []RegistrationRequestDTOResponse

func (h *RegistrationRequestsHTTPHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := core_http_request.GetLimitOffsetQueryParams(c)
	if err != nil {
		c.Error(err).SetMeta("failed to get 'limit'/'offset' query params")
		return
	}
	requests, err := h.registrationService.GetAll(ctx, limit, offset)
	if err != nil {
		c.Error(err).SetMeta("failed to get all requests")
		return
	}
	response := GetAllRequestsResponse(RegistrationRequestDTOFromDomains(requests))
	c.JSON(http.StatusOK, response)
}
