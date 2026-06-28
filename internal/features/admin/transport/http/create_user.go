package admin_transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
	admin_contracts "github.com/qandoni/keeneyePractice/internal/features/admin/contracts"
)

type CreateUserRequest struct {
	Login       string `json:"login" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Role        string `json:"role" binding:"required"`
	FIO         string `json:"fio"`
	PhoneNumber string `json:"phone_number"`
	GroupID     *int   `json:"group_id"`
}

type CreateUserResponse AdminDTOResponse

func (h *AdminHTTPHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var request CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to decode request",
		)
		return
	}
	command, err := commandFromRequest(request)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"invalid request",
		)
		return
	}

	userDomain, err := h.adminService.CreateUser(ctx, admin_contracts.CreateUserCommand(command))
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to create user",
		)
		return
	}
	response := CreateUserResponse(AdminDTOFromDomain(userDomain))
	c.JSON(http.StatusCreated, response)
}

func commandFromRequest(
	r CreateUserRequest,
) (admin_contracts.CreateUserCommand, error) {
	return admin_contracts.CreateUserCommand{
		Login:       r.Login,
		Password:    r.Password,
		Role:        enum.Role(r.Role),
		FIO:         r.FIO,
		PhoneNumber: r.PhoneNumber,
		GroupID:     r.GroupID,
	}, nil
}
