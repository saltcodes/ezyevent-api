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

// ListEvents func gets all existing events.
// @Description Get List of all events.
// @Summary get all existing events
// @Tags Events
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=[]model.Event} "desc"
// @Router /events [get]
func ListEvents(c *fiber.Ctx) error {
	var eventList []model.Event

	//Get address of the event list
	if err := queries.ListEvent(&eventList); err != nil {
		return util.CreateResponseMessage(c, util.InternalError, err.Error())
	}
	return util.CreateResponseMessage(c, util.Success, eventList)
}

// CreateEvent func create event in the database by checking with
// @Description create event in the database
// @Summary create an event
// @Tags Events
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.Event} "desc"
// @Param id path string true "Event ID"
// @Param request body model.Event true "query params"
// @Router /events/{id} [post]
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

// GetEvent func gets events by given ID or 404 error.
// @Description Get Event by given ID.
// @Summary get event by given ID
// @Tags Events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} model.ResponseObject{data=model.Event} "desc"
// @Router /events/{id} [get]
func GetEvent(c *fiber.Ctx) error {
	var event model.Event

	if err := queries.GetEvent(c.Params("id"), &event); err != nil {
		return util.CreateErrorResponseCode(c, err.Error())
	}

	return util.CreateResponseMessage(c, util.Success, event)
}

// UpdateEvent func updates events in the database by checking with id
// @Description updates events in the database by checking with id
// @Summary update event by id
// @Tags Events
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.Event} "desc"
// @Param id path string true "Event ID"
// @Param request body model.Event true "query params"
// @Router /events/{id} [put]
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
// @Description find events using lng,lat and radius
// @Tags Events
// @Summary query events bases and location
// @Accept json
// @Produce json
// @Param radius query string false "enter radius for event"
// @Param lat query string false "enter starting latitude"
// @Param lng query string false "enter starting  longitude"
// @Router /events/find [get]
// @Success 200 {object} model.ResponseObject{data=[]model.Event} "desc"
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

// DeleteEvent func delete events in the database by checking with id
// @Description delete events in the database by checking with id
// @Summary delete event by id
// @Tags Events
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseObject{data=model.Event} "desc"
// @Param id path string true "Event ID"
// @Param request body model.Event true "query params"
// @Router /events/{id} [delete]
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
