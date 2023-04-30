package handler

import (
	"context"
	"ezyevent-api/internal/model"
	"ezyevent-api/internal/proto"
	"ezyevent-api/internal/util"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var conn *grpc.ClientConn

func ListEvents(c *fiber.Ctx) error {
	var eventList []model.Event

	if err := (*db).Find(&eventList).Error; err != nil {
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

	if err := (*db).Create(event).Error; err != nil {
		return util.CreateResponseMessage(c, 500, err.Error(), nil)
	}

	//Extract Location object from struct
	location := &proto.LocationObject{
		Id: event.Id,
		Location: &proto.Location{
			Type:        "Type",
			Coordinates: []float32{float32(event.Lat), float32(event.Lng)},
		},
	}

	//Initiate gRPC client Connection
	con, err := grpc.Dial("localhost:8181", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf(err.Error())
	}

	//Call needed service from generated proto file with New_ServiceName
	client := proto.NewLocationDataServiceClient(con)
	_, err = client.LocationData(context.Background(), location)

	return util.CreateResponseMessage(c, 200, util.Success, event)
}
