package queries

import (
	"ezyevent-api/internal/database"
	"ezyevent-api/internal/model"
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

// FindEventsWithin accepts ids then return their respective event details
func FindEventsWithin(ids []string, eventList *[]model.Event) error {
	return (*db).Where("id IN ?", ids).Find(&eventList).Error
}

// UpdateEvent update events
func UpdateEvent(event *model.Event) error {
	return (*db).Save(event).Error
}

// DeleteEvent delete event
func DeleteEvent(id string) error {
	var event model.Event

	err := (*db).First(&event, "id=?", id).Error
	if event.Id == "" {
		return err
	}
	return (*db).Delete(&event).Error
}
