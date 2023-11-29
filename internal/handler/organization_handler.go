package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/queries"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

func CreateOrganization(c *fiber.Ctx) error {
	organization := new(model.Organization)

	//Parse Json Object to Struct
	if err := c.BodyParser(organization); err != nil {
		return util.CreateResponseMessage(c, util.InternalError, err.Error())
	}

	if err := queries.CreateOrganization(organization); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, organization)
}

func GetOrganization(c *fiber.Ctx) error {
	organization := new(model.Organization)

	if err := queries.GetOrganization(c.Params("id"), organization); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, organization)
}

func ListOrganization(c *fiber.Ctx) error {
	var organizations []model.Organization
	//var d = database.DBConn

	if err := queries.ListOrganization(&organizations); err != nil {
		return util.CreateResponseMessage(c, util.Success, nil)
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Message: "success",
		Status:  200,
	}, organizations)
}

func UpdateOrganization(c *fiber.Ctx) error {
	organization := new(model.Organization)

	if err := c.BodyParser(organization); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	if err := queries.UpdateOrganization(c.Params("id"), organization); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Message: "organization updated",
		Status:  201,
	}, organization)
}

func DeleteOrganization(c *fiber.Ctx) error {
	if err := queries.DeleteOrganization(c.Params("id")); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, util.Success, "organization deleted")
}
