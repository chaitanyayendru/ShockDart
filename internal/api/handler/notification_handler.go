package handler

import (
	"shockdart/internal/service"

	"github.com/gofiber/fiber/v2"
)

type NotificationHandler struct {
	service service.NotificationService
}

func NewNotificationHandler(s service.NotificationService) *NotificationHandler {
	return &NotificationHandler{service: s}
}

func (h *NotificationHandler) SendNotification(c *fiber.Ctx) error {
	var req service.NotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := h.service.SendNotification(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "Notification sent"})
}
