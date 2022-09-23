package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type EventType struct {
	Id   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             ` json:"name"`
}
