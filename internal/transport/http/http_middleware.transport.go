package http

import (
	"net/http"
	"time"

	"github.com/api-monolith-template/internal/model/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func loggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now().UTC()
		ctx.Next()
		endTime := time.Now().UTC()
		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()

		logrus.WithFields(logrus.Fields{
			"method":   reqMethod,
			"uri":      reqUri,
			"status":   statusCode,
			"latency":  latencyTime,
			"clientIP": clientIP,
		}).Info("Incoming request")

		ctx.Next()
	}
}

func customPanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Log the error
				c.Error(r.(error))

				// Respond with custom error message
				c.JSON(http.StatusInternalServerError, response.BaseResponse{
					Message: "internal server error",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
