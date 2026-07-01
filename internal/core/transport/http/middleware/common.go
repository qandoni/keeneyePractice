package core_http_middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	core_auth "github.com/qandoni/keeneyePractice/internal/core/auth"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
	core_logger "github.com/qandoni/keeneyePractice/internal/core/logger"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"

	"go.uber.org/zap"
)

const (
	requestIDHeader = "X-Request-ID"
	loggerKey       = "logger"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.GetHeader(requestIDHeader)
		if requestID == "" {
			requestID = uuid.NewString()
		}
		ctx.Writer.Header().Set(requestIDHeader, requestID)
		ctx.Request.Header.Set(requestIDHeader, requestID)

		ctx.Set(requestIDHeader, requestID)
		ctx.Next()
	}
}

func Logger(log *core_logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(requestIDHeader)

		requestLogger := log.With(
			zap.String("requestID", requestID),
			zap.String("url", c.Request.URL.String()),
		)

		c.Set(loggerKey, requestLogger)

		ctx := core_logger.ToContext(
			c.Request.Context(),
			requestLogger,
		)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := core_logger.FromContext(c.Request.Context())

		before := time.Now()

		log.Debug(
			">>> incoming HTTP request",
			zap.String("http_method", c.Request.Method),
			zap.Time("time", before.UTC()),
		)

		c.Next()

		log.Debug(
			"<<< done HTTP request",
			zap.Int("status_code", c.Writer.Status()),
			zap.Duration("latence", time.Since(before)),
		)
	}
}

type TokenParser interface {
	ParseAccessToken(
		token string,
	) (core_auth.AuthInfo, error)
}

func JWT(
	parser TokenParser,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			core_http_response.RespondError(
				c,
				core_errors.ErrUnauthorized,
				"missing authorization header",
			)
			c.Abort()
			return
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			core_http_response.RespondError(
				c,
				core_errors.ErrUnauthorized,
				"invalid authorizaiton header",
			)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, prefix)

		authInfo, err := parser.ParseAccessToken(
			token,
		)
		if err != nil {
			core_http_response.RespondError(
				c,
				err,
				"invalid token",
			)
			c.Abort()
			return
		}
		ctx := core_auth.WithAuthInfo(
			c.Request.Context(),
			authInfo,
		)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func Role(roles ...enum.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		authInfo, ok := core_auth.AuthInfoFromContext(c.Request.Context())
		if !ok {
			core_http_response.RespondError(
				c,
				core_errors.ErrUnauthorized,
				"authentication information not found",
			)
			c.Abort()
			return
		}

		for _, role := range roles {
			if authInfo.Role == role {
				c.Next()
				return
			}
		}

		core_http_response.RespondError(
			c,
			core_errors.ErrAccessForbidden,
			"access denied",
		)
		c.Abort()
		return
	}
}
