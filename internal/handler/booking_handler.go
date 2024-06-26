package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/queries"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

// CreateBooking func create booking in the database by checking with
// @Description create booking in the database
// @Summary create a booking
// @Tags Booking
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.Booking} "desc"
// @Param id path string true "Booking ID"
// @Param request body model.Booking true "query params"
// @Router /bookings/{id} [post]
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

// GetBooking func gets bookings by given ID or 404 error.
// @Description Get Booking by given ID.
// @Summary get booking by given ID
// @Tags Booking
// @Accept json
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} model.ResponseObject{data=model.Booking} "desc"
// @Router /bookings/{id} [get]
func GetBooking(c *fiber.Ctx) error {
	booking := new(model.Booking)

	if err := queries.GetBooking(c.Params("id"), booking); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, booking)
}

// ListBooking func gets all existing bookings.
// @Description Get List of all bookings.
// @Summary get all existing bookings
// @Tags Booking
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=[]model.Booking} "desc"
// @Router /bookings [get]
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

// UpdateBooking func updates bookings in the database by checking with id
// @Description updates bookings in the database by checking with id
// @Summary update booking by id
// @Tags Booking
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.Booking} "desc"
// @Param id path string true "Booking ID"
// @Param request body model.Booking true "query params"
// @Router /bookings/{id} [put]
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

// DeleteBooking func delete bookings in the database by checking with id
// @Description delete bookings in the database by checking with id
// @Summary delete booking by id
// @Tags Booking
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.Booking} "desc"
// @Param id path string true "Booking ID"
// @Param request body model.Booking true "query params"
// @Router /bookings/{id} [delete]
func DeleteBooking(c *fiber.Ctx) error {
	if err := queries.DeleteBooking(c.Params("id")); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, util.Success, "booking deleted")
}
