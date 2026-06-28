package admin_transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

type CreateGroupRequest struct {
	Name string `json:"name" binding:"required"`
}

func (h *AdminHTTPHandler) CreateGroup(c *gin.Context) {
	ctx := c.Request.Context()

	var req CreateGroupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		core_http_response.RespondError(c, err, "invalid request")
		return
	}

	group, err := h.adminService.CreateGroup(ctx, req.Name)
	if err != nil {
		core_http_response.RespondError(c, err, "failed to create group")
		return
	}

	c.JSON(http.StatusCreated, group)
}
