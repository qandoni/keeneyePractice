package users_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
)

type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()

	userID, err := core_http_request.GetIntPathValue(c, "id")

	if err != nil {
		c.Error(err).SetMeta("failed to get 'userID' path value")
		return
	}
	userDomain, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		c.Error(err).SetMeta("failed to get user")
		return
	}
	response := GetUserResponse(userDTOFromDomain(userDomain))
	c.JSON(http.StatusOK, response)
}
