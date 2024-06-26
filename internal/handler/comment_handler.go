package handler

import (
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
)

// ListComments func gets all existing comments.
// @Description Get List of all comments.
// @Summary get all existing comments
// @Tags Comments
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=[]model.Comment} "desc"
// @Router /comments [get]
func ListComments(c *fiber.Ctx) error {
	var comments []model.Comment

	if err := (*db).Preload("User").Find(&comments).Error; err != nil {
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
