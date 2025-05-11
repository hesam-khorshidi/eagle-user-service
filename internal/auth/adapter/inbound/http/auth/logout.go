package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/adapter/inbound/http"
)

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (c *Controller) Logout(ctx *fiber.Ctx) error {
	var request LogoutRequest
	if err := ctx.BodyParser(&request); err != nil {
		return http.BadRequest(ctx, err)
	}

	err := c.authService.Logout(ctx.UserContext(), request.RefreshToken)
	if err != nil {
		return http.DetermineErrorCode(ctx, err)
	}
	return http.Success(ctx, nil)
}
