package model

type Location struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type LocationObject struct {
	Id       string   `json:"id"`
	Location Location `json:"location"`
}
