package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

func ListCategories(c *fiber.Ctx) error {
	var categories []model.Category

	if err := (*db).Find(&categories).Error; err != nil {
		return util.CreateResponseMessage(c, model.StatusCode{
			Status:  500,
			Message: err.Error(),
		}, nil)
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Status:  200,
		Message: "success",
	}, categories)
}

func CreateCategory(c *fiber.Ctx) error {
	category := new(model.Category)

	if err := c.BodyParser(category); err != nil {
		return util.CreateResponseMessage(c, model.StatusCode{
			Status:  500,
			Message: err.Error(),
		}, nil)
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Status:  200,
		Message: "success",
	}, category)
}
