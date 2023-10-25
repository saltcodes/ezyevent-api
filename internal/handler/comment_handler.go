package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

func ListComments(c *fiber.Ctx) error {
	var comments []model.Comment

	if err := (*db).Find(&comments).Error; err != nil {
		return util.CreateResponseMessage(c, model.StatusCode{
			Status:  500,
			Message: err.Error(),
		}, nil)
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Status:  200,
		Message: "success",
	}, comments)
}
