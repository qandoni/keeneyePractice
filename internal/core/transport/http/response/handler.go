package core_http_response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	core_logger "github.com/qandoni/keeneyePractice/internal/core/logger"

	"go.uber.org/zap"
)

func RespondError(
	c *gin.Context,
	err error,
	msg string,
) {

	log := core_logger.FromContext(c.Request.Context())

	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {

	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = log.Warn

	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = log.Warn

	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = log.Debug
	case errors.Is(err, core_errors.ErrAccessForbidden):
		statusCode = http.StatusForbidden
		logFunc = log.Debug
	default:
		statusCode = http.StatusInternalServerError
		logFunc = log.Error
	}

	logFunc(
		msg,
		zap.Error(err),
	)

	c.JSON(
		statusCode,
		ErrorResponse{
			Error:   err.Error(),
			Message: msg,
		},
	)
}

func RespondNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
