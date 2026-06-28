package users_http_transport

import (
	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	userID, err := core_http_request.GetIntPathValue(c, "id")
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get int path value",
		)
		return
	}
	if err := h.usersService.DeleteUser(ctx, userID); err != nil {
		core_http_response.RespondError(c, err, "failed to delete user")
		return
	}
	core_http_response.RespondNoContent(c)

}
