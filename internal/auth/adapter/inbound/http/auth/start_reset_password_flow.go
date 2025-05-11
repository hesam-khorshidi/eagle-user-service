package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/adapter/inbound/http"
)

func (c *Controller) StartResetPasswordFlow(ctx *fiber.Ctx) error {
	return http.NoContent(ctx)
}
