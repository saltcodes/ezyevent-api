package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/queries"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

// CreateSchedule func create schedule in the database by checking with
// @Description create schedule in the database
// @Summary create a schedule
// @Tags Schedules
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.Schedule} "desc"
// @Param id path string true "Schedule ID"
// @Param request body model.Schedule true "query params"
// @Router /schedules/{id} [post]
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

// GetSchedule func gets schedules by given ID or 404 error.
// @Description Get Schedule by given ID.
// @Summary get schedule by given ID
// @Tags Schedules
// @Accept json
// @Produce json
// @Param id path string true "Schedule ID"
// @Success 200 {object} model.ResponseObject{data=model.Schedule} "desc"
// @Router /schedules/{id} [get]
func GetSchedule(c *fiber.Ctx) error {
	schedule := new(model.Schedule)

	if err := queries.GetSchedule(c.Params("id"), schedule); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, schedule)
}

// ListSchedule func gets all existing schedules.
// @Description Get List of all schedules.
// @Summary get all existing schedules
// @Tags Schedules
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=[]model.Schedule} "desc"
// @Router /schedules [get]
func ListSchedule(c *fiber.Ctx) error {
	var schedules []model.Schedule
	//var d = database.DBConn

	if err := queries.ListSchedule(&schedules); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Message: "success",
		Status:  200,
	}, schedules)
}

// UpdateSchedule func updates schedules in the database by checking with id
// @Description updates schedules in the database by checking with id
// @Summary update schedule by id
// @Tags Schedules
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.Schedule} "desc"
// @Param id path string true "Schedule ID"
// @Param request body model.Schedule true "query params"
// @Router /schedules/{id} [put]
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

// DeleteEventSchedule func delete schedules in the database by checking with id
// @Description delete schedules in the database by checking with id
// @Summary delete schedule by id
// @Tags Schedules
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.Schedule} "desc"
// @Param id path string true "Schedule ID"
// @Param request body model.Schedule true "query params"
// @Router /schedules/{id} [delete]
func DeleteEventSchedule(c *fiber.Ctx) error {
	if err := queries.DeleteSchedule(c.Params("id")); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, util.Success, "schedule deleted")
}
