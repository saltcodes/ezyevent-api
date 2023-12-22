package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/queries"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

func CreateSchedule(c *fiber.Ctx) error {
	schedule := new(model.Schedule)

	//Parse Json Object to Struct
	if err := c.BodyParser(schedule); err != nil {
		return util.CreateResponseMessage(c, util.InternalError, err.Error())
	}

	if err := queries.CreateSchedule(schedule); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, schedule)
}

func GetSchedule(c *fiber.Ctx) error {
	schedule := new(model.Schedule)

	if err := queries.GetSchedule(c.Params("id"), schedule); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, schedule)
}

func ListSchedule(c *fiber.Ctx) error {
	var schedules []model.Schedule
	//var d = database.DBConn

	if err := queries.ListSchedule(&schedules); err != nil {
		return util.CreateResponseMessage(c, util.Success, nil)
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Message: "success",
		Status:  200,
	}, schedules)
}

func UpdateSchedule(c *fiber.Ctx) error {
	schedule := new(model.Schedule)

	if err := c.BodyParser(schedule); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	if err := queries.UpdateSchedule(c.Params("id"), schedule); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Message: "schedule updated",
		Status:  201,
	}, schedule)
}

func DeleteEventSchedule(c *fiber.Ctx) error {
	if err := queries.DeleteSchedule(c.Params("id")); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, util.Success, "schedule deleted")
}
