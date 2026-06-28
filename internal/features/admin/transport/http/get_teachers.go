package admin_transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
)

type GetTeachersRequest struct {
	Limit  *int `form:"limit"`
	Offset *int `form:"offset"`
}

type TeacherDTO struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	FIO     string `json:"fio"`
	Version int    `json:"version"`
}

type GetTeachersResponse struct {
	Teachers []TeacherDTO `json:"teachers"`
}

func (h *AdminHTTPHandler) GetTeachers(c *gin.Context) {
	ctx := c.Request.Context()

	var req GetTeachersRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to parse query params",
		)
		return
	}

	teachers, err := h.adminService.GetTeachers(ctx, req.Limit, req.Offset)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get teachers",
		)
		return
	}

	resp := GetTeachersResponse{
		Teachers: make([]TeacherDTO, 0, len(teachers)),
	}

	for _, t := range teachers {
		resp.Teachers = append(resp.Teachers, TeacherDTO{
			ID:      t.ID,
			UserID:  t.UserID,
			FIO:     t.FIO,
			Version: t.Version,
		})
	}

	c.JSON(http.StatusOK, resp)
}
