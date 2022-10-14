package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

func ListComments(c *fiber.Ctx) error {
	var comments []model.Comment

	if err := (*db).Find(&comments).Error; err != nil {
		return util.CreateResponseMessage(c, 500, "internal server error", err.Error())
	}

	return util.CreateResponseMessage(c, 200, "success", comments)
}
