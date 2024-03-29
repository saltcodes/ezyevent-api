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
		return util.CreateResponseMessage(c, util.InternalError, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, eventList)
}

func CreateEvent(c *fiber.Ctx) error {
	event := new(model.Event)

	//Parse Json Object to Struct
	if err := c.BodyParser(event); err != nil {
		return util.CreateResponseMessage(c, util.InternalError, err.Error())
	}

	if err := queries.CreateEvent(event); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
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
	return util.CreateResponseMessage(c, util.Success, event)
}

func GetEvent(c *fiber.Ctx) error {
	var event model.Event

	if err := queries.GetEvent(c.Params("id"), &event); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, util.Success, event)
}

func UpdateEvent(c *fiber.Ctx) error {

	event := new(model.Event)

	//Parse Json Object to Struct
	if err := c.BodyParser(event); err != nil {
		return util.CreateResponseMessage(c, util.InternalError, err.Error())
	}

	if err := queries.UpdateEvent(c.Params("id"), event); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, model.StatusCode{
		Message: "event updated",
		Status:  201,
	}, event)
}

// FindEvents Find Events uses Post instead of GET to ensure correctness and also too lazy to be parsing params
func FindEvents(c *fiber.Ctx) error {
	var eventList []model.Event
	var eventIds []string

	//GRPC calls returns the whole
	eventsIdList, err := locationServerClient.FindEventsWithin(context.Background(), &proto.LocationQueryObject{
		Radius:    util.ConvertStringToInt32(c.Query("radius")),
		Latitude:  util.ConvertStringToFloat64(c.Query("lat")),
		Longitude: util.ConvertStringToFloat64(c.Query("lng")),
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, eventObject := range eventsIdList.GetEvents() {
		eventIds = append(eventIds, eventObject.Id)
	}

	if err := queries.FindEventsWithin(eventIds, &eventList); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, util.Success, eventList)
}

func DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := queries.Delete(id); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	if _, err := locationServerClient.DeleteEvent(context.Background(), &proto.EventId{
		Id: id,
	}); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, util.Success, "Event deleted")
}
