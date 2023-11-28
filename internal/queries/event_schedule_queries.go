package queries

import (
	"ezyevent-api/internal/model"
	"fmt"
	"strings"
)

// CreateScheduledEvent takes a model struct and saves it
func CreateScheduledEvent(eventSchedule *model.EventSchedule) error {
	return (*db).Create(eventSchedule).Error
}

// ListScheduledEvent Queries the eventSchedules tab;e and stored them in eventSchedule list array
func ListScheduledEvent(eventScheduleList *[]model.EventSchedule) error {
	return (*db).Find(&eventScheduleList).Error
}

// GetScheduledEvent Queries and return a single eventSchedule
func GetScheduledEvent(id string, eventSchedule *model.EventSchedule) error {
	return (*db).First(&eventSchedule, "id=?", id).Error
}

// UpdateScheduledEvent update eventSchedules
func UpdateScheduledEvent(id string, eventSchedule *model.EventSchedule) error {

	if strings.Compare(id, eventSchedule.Id) != 0 {
		return fmt.Errorf("invalid id")
	}

	return (*db).Save(eventSchedule).Error
}

// DeleteScheduledEvent delete eventSchedule
func DeleteScheduledEvent(id string) error {
	var eventSchedule model.EventSchedule

	err := (*db).First(&eventSchedule, "id=?", id).Error
	if eventSchedule.Id == "" {
		return err
	}
	return (*db).Delete(&eventSchedule).Error
}
