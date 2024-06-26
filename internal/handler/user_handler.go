package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/queries"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

// GetUsers func gets all existing users.
// @Description Get List of all users.
// @Summary get all existing users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=[]model.User} "desc"
// @Router /users [get]
func GetUsers(c *fiber.Ctx) error {
	var users []model.User
	//var d = database.DBConn

	if err := queries.ListUser(&users); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Message: "success",
		Status:  200,
	}, users)
}

// CreateUser func create user in the database by checking with
// @Description create user in the database
// @Summary create a user
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.User} "desc"
// @Param id path string true "User ID"
// @Param request body model.User true "query params"
// @Router /users/{id} [post]
func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	//Parse Json Object to Struct
	if err := c.BodyParser(user); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	if err := queries.CreateUser(user); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, user)
}

// GetUser func gets users by given ID or 404 error.
// @Description Get User by given ID.
// @Summary get user by given ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.ResponseObject{data=model.User} "desc"
// @Router /users/{id} [get]
func GetUser(c *fiber.Ctx) error {

	user := new(model.User)

	if err := queries.GetUser(c.Params("id"), user); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, user)
}

// DeleteUser func delete users in the database by checking with id
// @Description delete users in the database by checking with id
// @Summary delete user by id
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=string} "desc"
// @Param id path string true "User ID"
// @Param request body model.User true "query params"
// @Router /users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	if err := queries.DeleteUser(c.Params("id")); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, util.Success, "user deleted")
}

// UpdateUser func updates users in the database by checking with id
// @Description updates users in the database by checking with id
// @Summary update user by id
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.User} "desc"
// @Param id path string true "User ID"
// @Param request body model.User true "query params"
// @Router /users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {

	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	if err := queries.UpdateUser(c.Params("id"), user); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Message: "user updated",
		Status:  201,
	}, user)
}
