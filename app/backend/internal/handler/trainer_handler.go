package handler

import (
	"context"
	"strconv"

	"github.com/PawornpratKongdaeng/FitConnect/internal/domain"
	"github.com/PawornpratKongdaeng/FitConnect/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type TrainerHandler struct{ UC *usecase.TrainerUC }

type createTrainerReq struct {
	UserID      string   `json:"userId"`
	Bio         string   `json:"bio"`
	Specialties []string `json:"specialties"`
	Rate        int      `json:"rate"`
	CertStatus  string   `json:"certStatus"`
	RegionCode  string   `json:"regionCode"`
}

func (h *TrainerHandler) Create(c *fiber.Ctx) error {
	var req createTrainerReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	t := domain.Trainer{UserID: req.UserID, Bio: req.Bio, Specialties: req.Specialties, Rate: req.Rate, CertStatus: req.CertStatus, RegionCode: req.RegionCode}
	if err := h.UC.Create(context.Background(), t); err != nil {
		return fiber.NewError(400, err.Error())
	}
	return c.SendStatus(201)
}

func (h *TrainerHandler) ListByRegion(c *fiber.Ctx) error {
	region := c.Query("region")
	limitStr := c.Query("limit", "20")
	l, _ := strconv.ParseInt(limitStr, 10, 64)
	out, err := h.UC.ListByRegion(context.Background(), region, l)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	return c.JSON(out)
}
