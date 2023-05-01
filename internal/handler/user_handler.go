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
		return util.CreateResponseMessage(c, 500, "internal server error", nil)
	}

	return util.CreateResponseMessage(c, 200, "success", users)
}
