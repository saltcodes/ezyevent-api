package queries

import (
	"ezyevent-api/internal/model"
	"fmt"
	"strings"
)

// Create takes a model struct and saves it
func Create(data any) error {
	return (*db).Create(data).Error
}

// List queries the users tab;e and stored them in user list array
func List(data *[]any) error {
	return (*db).Find(&data).Error
}

// Select query and return a single object
func Select(id string, user *any) error {
	return (*db).First(&user, "id=?", id).Error
}

// Update updates data
func Update(id string, paramId string, data *any) error {

	if strings.Compare(id, paramId) != 0 {
		return fmt.Errorf("invalid id")
	}

	return (*db).Save(data).Error
}

// Delete data
func Delete(id string) error {
	var data any

	if err := (*db).First(&data, "id=?", id).Error; err != nil {
		return err
	}

	return (*db).Delete(&data).Error
}

// ListDataWithPreload  queries data with preload configuration
func ListDataWithPreload(preload string, data *any) error {
	return (*db).Preload(preload).Find(&data).Error
}

// FindEventsWithin accepts ids then return their respective event details
func FindEventsWithin(ids []string, eventList *[]model.Event) error {
	return (*db).Where("id IN ?", ids).Find(&eventList).Error
}
