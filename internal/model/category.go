package model

type Category struct {
	Id      string `json:"id"`
	Name    string ` json:"name,omitempty"`
	IconUrl string `json:"iconUrl,omitempty"`
}
