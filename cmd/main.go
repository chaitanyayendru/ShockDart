package main

import (
	"log"
	"shockdart/config"
	"shockdart/internal/api"
	"shockdart/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config := config.LoadConfig()

	logger.Initialize()

	api.SetupRoutes(app)

	if err := app.Listen(config.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
