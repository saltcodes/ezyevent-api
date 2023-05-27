package handler

import (
	"context"
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/proto"
	"ezyevent-api/internal/queries"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
	"log"
)

func ListEvents(c *fiber.Ctx) error {
	var eventList []model.Event

	//Get address of the event list
	if err := queries.ListEvent(&eventList); err != nil {
		return util.CreateResponseMessage(c, fiber.StatusInternalServerError, util.InternalError, err.Error())
	}
	return util.CreateResponseMessage(c, 200, util.Success, eventList)
}

func CreateEvent(c *fiber.Ctx) error {
	event := new(model.Event)

	//Parse Json Object to Struct
	if err := c.BodyParser(event); err != nil {
		return util.CreateResponseMessage(c, 500, util.InternalError, err.Error())
	}

	if err := queries.CreateEvent(event); err != nil {
		return util.CreateResponseMessage(c, 500, err.Error(), nil)
	}

	//Extract Location object from struct
	location := &proto.LocationObject{
		Id: event.Id,
		Location: &proto.Location{
			Type:        "Point",
			Coordinates: []float32{float32(event.Lat), float32(event.Lng)},
		},
	}

	if grpcErr != nil {
		log.Fatalf(grpcErr.Error())
	}

	//Send location data to location RPC for geo hashing
	client := proto.NewLocationDataServiceClient(con)
	_, grpcErr = client.LocationData(context.Background(), location)
	return util.CreateResponseMessage(c, 200, util.Success, event)
}
