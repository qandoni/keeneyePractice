package users_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
)

type GetUsersResponse []UserDTOResponse

func (h *UsersHTTPHandler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := core_http_request.GetLimitOffsetQueryParams(c)
	if err != nil {
		c.Error(err).SetMeta("failed to get 'limit'/'offset' query params")
		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		c.Error(err).SetMeta("failed to get users")
		return
	}

	response := GetUsersResponse(usersDTOFromDomains(userDomains))
	c.JSON(http.StatusOK, response)
}
