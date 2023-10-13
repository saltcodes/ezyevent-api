package util

import (
	"ezyevent-api/internal/model"
	"github.com/gofiber/fiber/v2"
)

func CreateResponseMessage(c *fiber.Ctx, status model.StatusCode, data any) error {
	return c.Status(status.Status).JSON(&fiber.Map{
		"code":    status.Status,
		"message": status.Message,
		"data":    data,
	})
}

func CreateErrorResponseCode(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
		"code":    fiber.StatusInternalServerError,
		"message": message,
	})
}
