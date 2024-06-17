package auth

import "github.com/api-monolith-template/internal/model/contract"

type Controller struct {
	authService contract.AuthService
}

func NewController() *Controller {
	return new(Controller)
}

func (c *Controller) WithAuthService(svc contract.AuthService) *Controller {
	c.authService = svc
	return c
}
