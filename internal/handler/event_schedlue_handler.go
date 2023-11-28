package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/queries"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

func CreateEventSchedule(c *fiber.Ctx) error {
	eventSchedule := new(model.EventSchedule)

	//Parse Json Object to Struct
	if err := c.BodyParser(eventSchedule); err != nil {
		return util.CreateResponseMessage(c, util.InternalError, err.Error())
	}

	if err := queries.CreateScheduledEvent(eventSchedule); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, eventSchedule)
}

func GetEventSchedule(c *fiber.Ctx) error {
	eventSchedule := new(model.EventSchedule)

	if err := queries.GetScheduledEvent(c.Params("id"), eventSchedule); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, eventSchedule)
}

func ListEventSchedule(c *fiber.Ctx) error {
	var eventSchedules []model.EventSchedule
	//var d = database.DBConn

	if err := queries.ListScheduledEvent(&eventSchedules); err != nil {
		return util.CreateResponseMessage(c, util.Success, nil)
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Message: "success",
		Status:  200,
	}, eventSchedules)
}

func UpdateEventSchedule(c *fiber.Ctx) error {
	eventSchedule := new(model.EventSchedule)

	if err := c.BodyParser(eventSchedule); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	if err := queries.UpdateScheduledEvent(c.Params("id"), eventSchedule); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Message: "eventSchedule updated",
		Status:  201,
	}, eventSchedule)
}

func DeleteEventSchedule(c *fiber.Ctx) error {
	if err := queries.DeleteScheduledEvent(c.Params("id")); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, util.Success, "event deleted")
}
