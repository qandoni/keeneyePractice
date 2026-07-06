package registration_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	registration_contracts "github.com/qandoni/keeneyePractice/internal/features/registration_requests/contracts"
)

type CompleteRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

func (h *RegistrationRequestsHTTPHandler) Complete(c *gin.Context) {
	ctx := c.Request.Context()

	var req CompleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err).SetMeta("invalid request")
		return
	}

	err := h.service.Complete(ctx, registration_contracts.CompleteInput{
		Token:    req.Token,
		Password: req.Password,
	})

	if err != nil {
		c.Error(err).SetMeta("failed to complete registration")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "registration completed",
	})
}
