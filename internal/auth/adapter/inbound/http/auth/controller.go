package auth

import (
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/port/inbound"
	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/adapter/inbound/http"
)

type Controller struct {
	authService inbound.AuthorizationService
}

func Init(d http.Dependencies, authService inbound.AuthorizationService) Controller {
	c := Controller{
		authService: authService,
	}

	authGroup := d.Fiber.Group(d.Prefix + "/auth")
	authGroup.Post("/register", c.Register)
	authGroup.Post("/login", c.Login)
	authGroup.Post("/logout", c.Logout)
	authGroup.Post("/refresh", c.RefreshToken)
	authGroup.Post("/reset-password/start", c.StartResetPasswordFlow)
	authGroup.Post("/reset-password/do", c.ResetPassword)

	return c
}
