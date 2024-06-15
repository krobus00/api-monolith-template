package http

import (
	"time"

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
