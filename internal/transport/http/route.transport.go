package http

import (
	"github.com/api-monolith-template/internal/constant"
	"github.com/api-monolith-template/internal/model/response"
	"github.com/gin-gonic/gin"
)

func (t *Transport) InitRoute() {
	t.router.Use(t.middlewareController.CustomPanicHandler(), t.middlewareController.LoggingMiddleware())

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
	authRefreshToken := authGroup.Group("/refresh", t.middlewareController.AuthMiddleware(constant.RefreshTokenType))
	authRefreshToken.POST("/", t.authController.RefreshToken)

	authProtected := authGroup.Use(t.middlewareController.AuthMiddleware(constant.AccessTokenType))
	authProtected.GET("/info", t.authController.Info)
	authProtected.POST("/logout", t.authController.Logout)

	// handle route not found
	t.router.NoRoute(func(c *gin.Context) {
		resp := constant.ErrEndpointNotFound.ToResponse()
		c.JSON(resp.StatusCode, resp)
	})

}
