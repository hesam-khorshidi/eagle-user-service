package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/adapter/inbound/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c Controller) Login(ctx *fiber.Ctx) error {
	var request LoginRequest
	if err := ctx.BodyParser(&request); err != nil {
		return http.BadRequest(ctx, err)
	}
	token, err := c.authService.Login(ctx.UserContext(), request.Email, request.Password)
	if err != nil {
		return http.InternalError(ctx, err)
	}
	return http.Success(ctx, token)
}
