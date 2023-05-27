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

var locationServerClient = proto.NewLocationDataServiceClient(con)

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
			Coordinates: []float64{event.Lat, event.Lng},
		},
	}

	if grpcErr != nil {
		log.Fatalf(grpcErr.Error())
	}

	//Send location data to location RPC for geo hashing
	_, grpcErr = locationServerClient.LocationData(context.Background(), location)
	return util.CreateResponseMessage(c, 200, util.Success, event)
}

// FindEvents Find Events uses Post instead of GET to ensure correctness and also too lazy to be parsing params
func FindEvents(c *fiber.Ctx) error {
	locationQueryObject := new(proto.LocationQueryObject)
	var eventList []model.Event
	var eventIds []string

	//Parse location Query Object which includes radius,latitude and longitude.
	if err := c.BodyParser(locationQueryObject); err != nil {
		return util.CreateResponseMessage(c, 500, err.Error(), nil)
	}

	//GRPC calls returns the whole
	eventsIdList, err := locationServerClient.FindEventsWithin(context.Background(), locationQueryObject)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, eventObject := range eventsIdList.GetEvents() {
		eventIds = append(eventIds, eventObject.Id)
	}

	if err := queries.FindEventsWithin(eventIds, &eventList); err != nil {
		return util.CreateResponseMessage(c, 500, err.Error(), nil)
	}

	return util.CreateResponseMessage(c, 200, util.Success, eventList)
}
