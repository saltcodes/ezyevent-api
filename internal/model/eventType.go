package model

type EventType struct {
	Id      string `json:"id"`
	Name    string ` json:"name,omitempty"`
	IconUrl string `json:"iconUrl,omitempty"`
}
