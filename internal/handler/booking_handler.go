package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/queries"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

func CreateBooking(c *fiber.Ctx) error {
	booking := new(model.Booking)

	//Parse Json Object to Struct
	if err := c.BodyParser(booking); err != nil {
		return util.CreateResponseMessage(c, util.InternalError, err.Error())
	}

	if err := queries.CreateBooking(booking); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, booking)
}

func GetBooking(c *fiber.Ctx) error {
	booking := new(model.Booking)

	if err := queries.GetBooking(c.Params("id"), booking); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, booking)
}

func ListBooking(c *fiber.Ctx) error {
	var bookings []model.Booking
	//var d = database.DBConn

	if err := queries.ListBooking(&bookings); err != nil {
		return util.CreateResponseMessage(c, util.Success, nil)
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Message: "success",
		Status:  200,
	}, bookings)
}

func UpdateBooking(c *fiber.Ctx) error {
	booking := new(model.Booking)

	if err := c.BodyParser(booking); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	if err := queries.UpdateBooking(c.Params("id"), booking); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Message: "booking updated",
		Status:  201,
	}, booking)
}

func DeleteBooking(c *fiber.Ctx) error {
	if err := queries.DeleteBooking(c.Params("id")); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, util.Success, "booking deleted")
}
