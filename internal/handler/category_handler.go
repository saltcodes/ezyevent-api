package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

// ListCategories func gets all existing categories.
// @Description Get List of all categories.
// @Summary get all existing categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=[]model.Category} "desc"
// @Router /categories [get]
func ListCategories(c *fiber.Ctx) error {
	var categories []model.Category

	if err := (*db).Find(&categories).Error; err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Status:  200,
		Message: "success",
	}, categories)
}

// CreateCategory func create category in the database by checking with
// @Description create category in the database
// @Summary create a category
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.Category} "desc"
// @Param id path string true "Category ID"
// @Param request body model.Category true "query params"
// @Router /categories/{id} [post]
func CreateCategory(c *fiber.Ctx) error {
	category := new(model.Category)

	if err := c.BodyParser(category); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Status:  200,
		Message: "success",
	}, category)
}
