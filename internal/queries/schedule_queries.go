package queries

import (
	"ezyevent-api/internal/model"
	"fmt"
	"strings"
)

// CreateSchedule takes a model struct and saves it
func CreateSchedule(schedule *model.Schedule) error {
	return (*db).Create(schedule).Error
}

// ListSchedule Queries the eventSchedules tab;e and stored them in eventSchedule list array
func ListSchedule(scheduleList *[]model.Schedule) error {
	return (*db).Preload("Event").Find(&scheduleList).Error
}

// GetSchedule Queries and return a single eventSchedule
func GetSchedule(id string, schedule *model.Schedule) error {
	return (*db).Preload("Event").First(&schedule, "id=?", id).Error
}

// UpdateSchedule update eventSchedules
func UpdateSchedule(id string, schedule *model.Schedule) error {

	if strings.Compare(id, schedule.Id) != 0 {
		return fmt.Errorf("invalid id")
	}

	return (*db).Save(schedule).Error
}

// DeleteSchedule delete eventSchedule
func DeleteSchedule(id string) error {
	var eventSchedule model.Schedule

	err := (*db).First(&eventSchedule, "id=?", id).Error
	if eventSchedule.Id == "" {
		return err
	}
	return (*db).Delete(&eventSchedule).Error
}
