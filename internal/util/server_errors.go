package util

import "ezyevent-api/internal/model"

var InternalError = model.StatusCode{
	Status:  200,
	Message: "Internal server Error",
}

var Success = model.StatusCode{
	Status:  200,
	Message: "Success",
}
