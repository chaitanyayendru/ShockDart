package api

import (
	"shockdart/internal/api/handler"
	"shockdart/internal/repository"
	"shockdart/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	notificationRepo := repository.NewKafkaRepository() // Kafka for queuing
	notificationService := service.NewNotificationService(notificationRepo)
	notificationHandler := handler.NewNotificationHandler(notificationService)

	api := app.Group("/api")
	api.Post("/notification", notificationHandler.SendNotification)
}
