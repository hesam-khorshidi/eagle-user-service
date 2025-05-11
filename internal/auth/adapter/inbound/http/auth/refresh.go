package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/adapter/inbound/http"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

func (c *Controller) RefreshToken(ctx *fiber.Ctx) error {
	var request RefreshRequest
	if err := ctx.BodyParser(&request); err != nil {
		return http.BadRequest(ctx, err)
	}
	token, err := c.authService.RefreshToken(ctx.UserContext(), request.RefreshToken)
	if err != nil {
		return http.DetermineErrorCode(ctx, err)
	}
	return http.Success(ctx, RefreshResponse{AccessToken: token})
}
