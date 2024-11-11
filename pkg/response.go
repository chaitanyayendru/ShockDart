package pkg

import "github.com/gofiber/fiber/v2"

func JSON(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{"data": data})
}
