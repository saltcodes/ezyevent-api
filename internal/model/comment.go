package model

type Comment struct {
	ID      string `json:"id"`
	UserId  string `json:"user_id"`
	Message string `json:"message"`
	EventId string `json:"eventId"`
}
