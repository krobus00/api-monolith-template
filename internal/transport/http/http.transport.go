package http

import (
	"github.com/api-monolith-template/internal/transport/http/auth"
	"github.com/api-monolith-template/internal/transport/http/middleware"
	"github.com/gin-gonic/gin"
)

type Transport struct {
	router *gin.Engine

	authController       *auth.Controller
	middlewareController *middleware.Controller
}

func NewTransport() *Transport {
	return new(Transport)
}

func (t *Transport) WithGinEngine(r *gin.Engine) *Transport {
	t.router = r
	return t
}

func (t *Transport) WithAuthController(c *auth.Controller) *Transport {
	t.authController = c
	return t
}

func (t *Transport) WithMiddlewareController(c *middleware.Controller) *Transport {
	t.middlewareController = c
	return t
}
