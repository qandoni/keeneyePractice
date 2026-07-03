package users_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
)

func (h *UsersHTTPHandler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	userID, err := core_http_request.GetIntPathValue(c, "id")
	if err != nil {
		c.Error(err).SetMeta("failed to get int path value")
		return
	}
	if err := h.usersService.DeleteUser(ctx, userID); err != nil {
		c.Error(err).SetMeta("failed to delete user")
		return
	}
	c.Status(http.StatusNoContent)

}
