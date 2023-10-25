package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

func ListEventTypes(c *fiber.Ctx) error {
	var eventTypes []model.EventType

	if err := (*db).Find(&eventTypes).Error; err != nil {
		return util.CreateResponseMessage(c, model.StatusCode{
			Status:  500,
			Message: err.Error(),
		}, nil)
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Status:  200,
		Message: "success",
	}, eventTypes)
}

func CreateEventType(c *fiber.Ctx) error {
	eventType := new(model.EventType)

	if err := c.BodyParser(eventType); err != nil {
		return util.CreateResponseMessage(c, model.StatusCode{
			Status:  500,
			Message: err.Error(),
		}, nil)
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Status:  200,
		Message: "success",
	}, eventType)
}
