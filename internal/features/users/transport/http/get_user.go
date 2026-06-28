package users_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()

	userID, err := core_http_request.GetIntPathValue(c, "id")

	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get 'userID' path value",
		)
		return
	}
	userDomain, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get user",
		)
		return
	}
	response := GetUserResponse(userDTOFromDomain(userDomain))
	c.JSON(http.StatusOK, response)
}
