package util

import "github.com/gofiber/fiber/v2"

func CreateResponseMessage(c *fiber.Ctx, status int, message string, data any) error {
	return c.Status(status).JSON(&fiber.Map{
		"code":    status,
		"message": message,
		"data":    data,
	})
}
