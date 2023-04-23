package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/util"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ListEvents(c *fiber.Ctx) error {
	var eventList []model.Event

	if err := (*db).Find(&eventList).Error; err != nil {
		return util.CreateResponseMessage(c, fiber.StatusInternalServerError, util.InternalError, err.Error())
	}
	return util.CreateResponseMessage(c, 200, util.Success, eventList)
}

func CreateEvent(c *fiber.Ctx) error {
	event := new(model.Event)

	if err := c.BodyParser(event); err != nil {
		return util.CreateResponseMessage(c, 500, util.InternalError, err.Error())
	}

	if err := (*db).Create(event).Error; err != nil {
		return util.CreateResponseMessage(c, 500, err.Error(), nil)
	}

	location := model.LocationObject{Id: event.Id, Location: model.Location{Type: "Point", Coordinates: []float64{event.Lat, event.Lng}}}
	fmt.Print(location)

	return util.CreateResponseMessage(c, 200, util.Success, event)
}
