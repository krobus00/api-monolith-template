package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *Transport) InitRoute() {
	t.router.Use(loggingMiddleware())

	internalGroup := t.router.Group("/_internal")
	internalGroup.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
}
