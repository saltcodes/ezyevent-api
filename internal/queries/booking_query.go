package queries

import (
	"ezyevent-api/internal/model"
	"fmt"
	"strings"
)

// CreateBooking takes a model struct and saves it
func CreateBooking(booking *model.Booking) error {
	return (*db).Create(booking).Error
}

// ListBooking Queries the events tab;e and stored them in booking list array
func ListBooking(bookingList *[]model.Booking) error {
	return (*db).Preload("User").Preload("Schedule").Find(&bookingList).Error
}

// GetBooking Queries and return a single booking
func GetBooking(id string, booking *model.Booking) error {
	return (*db).Preload("User").Preload("Schedule").First(&booking, "id=?", id).Error
}

// UpdateBooking update events
func UpdateBooking(id string, booking *model.Booking) error {

	if strings.Compare(id, booking.ID) != 0 {
		return fmt.Errorf("invalid id")
	}

	return (*db).Save(booking).Error
}

// DeleteBooking delete booking
func DeleteBooking(id string) error {
	var booking model.Booking

	if err := (*db).First(&booking, "id=?", id).Error; err != nil {
		return err
	}

	return (*db).Delete(&booking).Error
}
