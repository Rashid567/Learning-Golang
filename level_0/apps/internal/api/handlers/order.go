package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/rashid567/learning-golang/level_0/apps/internal/models"
	"github.com/rashid567/learning-golang/level_0/apps/internal/utils"
)

func (s *Handler) GetOrder(c fiber.Ctx) error {
	orderUID := models.OrderUID(c.Params("order_uid"))
	order, err := s.orderService.GetOrder(orderUID)

	if err != nil {
		var app_err utils.AppError
		if errors.As(err, &app_err) {
			return c.Status(app_err.StatusCode).JSON(app_err.Detail)
		}
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(order)
}
