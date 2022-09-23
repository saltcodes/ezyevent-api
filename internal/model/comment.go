package model

type Comment struct {
	Id      string `bson:"_id,omitempty" json:"id"`
	UserId  string `json:"user_id"`
	Message string `json:"message"`
	EventId string `json:"eventId"`
}
