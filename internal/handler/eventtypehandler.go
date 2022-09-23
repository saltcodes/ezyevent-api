package handler

import (
	"context"
	"ezyevent-api/internal/database"
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

var db = database.Config()
var ctx = context.TODO()
var eventCollection = db.Collection("eventTypes")

func ListEventTypes(c *fiber.Ctx) error {
	var eventTypes []model.EventType
	cursor, err := eventCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &eventTypes); err != nil {
		return util.CreateResponseMessage(c, 500, "internal server error", nil)
	}

	return util.CreateResponseMessage(c, 200, "success", eventTypes)
}

func CreateEventType(c *fiber.Ctx) error {
	eventType := new(model.EventType)

	if err := c.BodyParser(eventType); err != nil {
		return util.CreateResponseMessage(c, 500, "internal server error", err.Error())
	}

	result, err := eventCollection.InsertOne(ctx, eventType)
	if err != nil {
		return util.CreateResponseMessage(c, 500, "internal server error", err.Error())
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		eventType.Id = oid.Hex()
	}

	return util.CreateResponseMessage(c, 200, "success", eventType)
}
