package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

func ListEventTypes(c *fiber.Ctx) error {
	var eventTypes []model.EventType

	if err := (*db).Find(&eventTypes).Error; err != nil {
		return util.CreateResponseMessage(c, 500, "internal server error", err.Error())
	}

	return util.CreateResponseMessage(c, 200, "success", eventTypes)
}

func CreateEventType(c *fiber.Ctx) error {
	eventType := new(model.EventType)

	if err := c.BodyParser(eventType); err != nil {
		return util.CreateResponseMessage(c, 500, "internal server error", err.Error())
	}

	return util.CreateResponseMessage(c, 200, "success", eventType)
}
