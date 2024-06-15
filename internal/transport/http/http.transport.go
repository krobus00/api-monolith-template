package http

import "github.com/gin-gonic/gin"

type Transport struct {
	router *gin.Engine
}

func NewTransport() *Transport {
	return new(Transport)
}

func (t *Transport) WithGinEngine(r *gin.Engine) *Transport {
	t.router = r
	return t
}
