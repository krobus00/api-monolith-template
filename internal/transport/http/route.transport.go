package http

import (
	"github.com/api-monolith-template/internal/config"
	"github.com/api-monolith-template/internal/model/response"
	"github.com/gin-gonic/gin"
)

func (t *Transport) InitRoute() {
	t.router.Use(customPanicHandler(), loggingMiddleware())

	internalGroup := t.router.Group("/_internal")
	internalGroup.GET("/healthz", func(c *gin.Context) {
		resp := response.NewResponseOK()
		c.JSON(resp.StatusCode, resp)
	})

	v1Group := t.router.Group("/v1")

	authGroup := v1Group.Group("/auth")
	authGroup.POST("/register", t.authController.Register)
	authGroup.POST("/login", t.authController.Login)

	authProtected := authGroup.Use(AuthMiddleware(config.Env.Token.AccessTokenSecret))
	authProtected.GET("/info", t.authController.Info)
}
