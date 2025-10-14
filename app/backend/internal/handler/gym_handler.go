package handler

import (
	"context"
	"strconv"

	"github.com/PawornpratKongdaeng/FitConnect/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type GymHandler struct{ UC *usecase.GymUC }

func (h *GymHandler) Near(c *fiber.Ctx) error {
	lat, _ := strconv.ParseFloat(c.Query("lat"), 64)
	lng, _ := strconv.ParseFloat(c.Query("lng"), 64)
	km, _ := strconv.ParseFloat(c.Query("km", "5"), 64)
	out, err := h.UC.Near(context.Background(), lat, lng, km)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	return c.JSON(out)
}
