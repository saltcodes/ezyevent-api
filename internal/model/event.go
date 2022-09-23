package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
}

type Event struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name"`
	Summary  string             `json:"summary"`
	Details  string             `json:"details"`
	Images   []string           `json:"images"`
	Location struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"location"`
	Price  float64 `json:"price"`
	Date   int     `json:"date"`
	TypeID string  `json:"type-id"`
}
