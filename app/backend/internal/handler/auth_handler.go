package handler

import (
	"context"

	"github.com/PawornpratKongdaeng/FitConnect/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct{ UC *usecase.AuthUC }

type registerReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req registerReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	u, err := h.UC.Register(context.Background(), usecase.RegisterInput{Name: req.Name, Email: req.Email, Password: req.Password})
	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	return c.JSON(u)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req loginReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	token, err := h.UC.Login(context.Background(), usecase.LoginInput{Email: req.Email, Password: req.Password})
	if err != nil {
		return fiber.NewError(401, err.Error())
	}
	return c.JSON(fiber.Map{"accessToken": token})
}
