package model

type EventType struct {
	Id      string `bson:"_id,omitempty" json:"id"`
	Name    string ` json:"name,omitempty"`
	IconUrl string `json:"iconUrl,omitempty"`
}
