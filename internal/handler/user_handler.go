package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/queries"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []model.User
	//var d = database.DBConn

	if err := queries.ListUser(&users); err != nil {
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

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	//Parse Json Object to Struct
	if err := c.BodyParser(user); err != nil {
		return util.CreateResponseMessage(c, util.InternalError, err.Error())
	}

	if err := queries.CreateUser(user); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, user)
}
