package queries

import (
	"ezyevent-api/internal/database"
	"ezyevent-api/internal/model"
	"fmt"
	"strings"
)

var db = &database.DBConn

// CreateEvent takes a model struct and saves it
func CreateEvent(event *model.Event) error {
	return (*db).Create(event).Error
}

// ListEvent Queries the events tab;e and stored them in event list array
func ListEvent(eventList *[]model.Event) error {
	return (*db).Preload("EventType").Find(&eventList).Error
}

// GetEvent Queries and return a single event
func GetEvent(id string, event *model.Event) error {
	return (*db).First(&event, "id=?", id).Error
}

// UpdateEvent update events
func UpdateEvent(id string, event *model.Event) error {

	if strings.Compare(id, event.Id) != 0 {
		return fmt.Errorf("invalid id")
	}

	return (*db).Save(event).Error
}

// DeleteEvent delete event
func DeleteEvent(id string) error {
	var event model.Event

	if err := (*db).First(&event, "id=?", id).Error; err != nil {
		return err
	}

	return (*db).Delete(&event).Error
}
