package queries

import (
	"ezyevent-api/internal/database"
	"ezyevent-api/internal/model"
)

var db = &database.DBConn

func CreateEvent(event *model.Event) error {
	return (*db).Create(event).Error
}

// ListEvent Queries the events tab;e and stored them in event list array
func ListEvent(eventList *[]model.Event) error {
	return (*db).Find(&eventList).Error
}
