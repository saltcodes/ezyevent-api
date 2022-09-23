package handler

import (
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

func CreateEvent(c *fiber.Ctx) error {

	return util.CreateResponseMessage(c, 202, "success", nil)
}
