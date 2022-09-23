package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id         primitive.ObjectID `bson:"id" json:"id"`
	FirstName  string             `bson:"firstName"`
	LastName   string             `bson:"lastName"`
	Email      string             `bson:"email"`
	Phone      string             `bson:"phone"`
	ProfileUrl string             `bson:"profileUrl"`
}
