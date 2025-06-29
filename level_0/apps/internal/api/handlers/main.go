package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rashid567/learning-golang/level_0/apps/internal/domain/order"
)

type Handler struct {
	orderService order.OrderService
}

func NewHandler(orderService order.OrderService) *Handler {
	return &Handler{
		orderService: orderService,
	}
}

func (h *Handler) Register(app *fiber.App) {
	order_group := app.Group("/order")

	order_group.Get("/:order_uid", h.GetOrder)
}
