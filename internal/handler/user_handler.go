package handler

import (
	"ezyevent-api/internal/database"
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []model.User
	//var d = database.DBConn

	if err := database.DBConn.Find(&users).Error; err != nil {
		return util.CreateResponseMessage(c, model.StatusCode{
			Message: "success",
			Status:  201,
		}, nil)
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Message: "success",
		Status:  200,
	}, users)
}
