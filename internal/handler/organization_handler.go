package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/queries"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

// CreateOrganization func create organization in the database by checking with
// @Description create organization in the database
// @Summary create an organization
// @Tags Organizations
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.Organization} "desc"
// @Param id path string true "Organization ID"
// @Param request body model.Organization true "query params"
// @Router /organizations/{id} [post]
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

// GetOrganization func gets organizations by given ID or 404 error.
// @Description Get Organization by given ID.
// @Summary get organization by given ID
// @Tags Organizations
// @Accept json
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} model.ResponseObject{data=model.Organization} "desc"
// @Router /organizations/{id} [get]
func GetOrganization(c *fiber.Ctx) error {
	organization := new(model.Organization)

	if err := queries.GetOrganization(c.Params("id"), organization); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, organization)
}

// ListOrganization func gets all existing organizations.
// @Description Get List of all organizations.
// @Summary get all existing organizations
// @Tags Organizations
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=[]model.Organization} "desc"
// @Router /organizations [get]
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

// UpdateOrganization func updates organizations in the database by checking with id
// @Description updates organizations in the database by checking with id
// @Summary update organization by id
// @Tags Organizations
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.Organization} "desc"
// @Param id path string true "Organization ID"
// @Param request body model.Organization true "query params"
// @Router /organizations/{id} [put]
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

// DeleteOrganization func delete organizations in the database by checking with id
// @Description delete organizations in the database by checking with id
// @Summary delete organization by id
// @Tags Organizations
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.Organization} "desc"
// @Param id path string true "Organization ID"
// @Param request body model.Organization true "query params"
// @Router /organizations/{id} [delete]
func DeleteOrganization(c *fiber.Ctx) error {
	if err := queries.DeleteOrganization(c.Params("id")); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, util.Success, "organization deleted")
}
