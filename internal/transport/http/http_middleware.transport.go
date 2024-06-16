package http

import (
	"net/http"
	"time"

	"github.com/api-monolith-template/internal/constant"
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
				if err, ok := r.(error); ok {
					logrus.Error(err)
				}
				c.JSON(http.StatusInternalServerError, constant.ErrInternalServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
