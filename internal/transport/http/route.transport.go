package http

import (
	"github.com/api-monolith-template/internal/config"
	"github.com/api-monolith-template/internal/constant"
	"github.com/api-monolith-template/internal/model/response"
	"github.com/gin-gonic/gin"
)

func (t *Transport) InitRoute() {
	t.router.Use(customPanicHandler(), loggingMiddleware())

	// handle health check
	internalGroup := t.router.Group("/_internal")
	internalGroup.GET("/healthz", func(c *gin.Context) {
		resp := response.NewResponseOK()
		c.JSON(resp.StatusCode, resp)
	})

	v1Group := t.router.Group("/v1")

	authGroup := v1Group.Group("/auth")
	authGroup.POST("/register", t.authController.Register)
	authGroup.POST("/login", t.authController.Login)
	authRefreshToken := authGroup.Group("/refresh", AuthMiddleware(config.Env.Token.RefreshTokenSecret))
	authRefreshToken.POST("/", t.authController.RefreshToken)

	authProtected := authGroup.Use(AuthMiddleware(config.Env.Token.AccessTokenSecret))
	authProtected.GET("/info", t.authController.Info)

	// handle route not found
	t.router.NoRoute(func(c *gin.Context) {
		resp := constant.ErrEndpointNotFound.ToResponse()
		c.JSON(resp.StatusCode, resp)
	})

}
