package model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Event struct {
	Id           string         `json:"id"`
	Name         string         `json:"name"`
	Summary      string         `json:"summary"`
	Details      string         `json:"details"`
	Images       pq.StringArray `json:"images" gorm:"type:string[]"`
	Banner       string         `json:"banner"`
	Lat          float64        `json:"lat"`
	Lng          float64        `json:"lng"`
	LocationName string         `json:"locationName"`
	Price        float64        `json:"price"`
	Date         int            `json:"date"`
	EventTypeID  string         `json:"eventTypeID" gorm:"event_type_id"`
	EventType    EventType      `json:"eventType" gorm:"foreignKey:EventTypeID"`
}

func (base *Event) BeforeCreate(db *gorm.DB) (err error) {
	uId, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	base.Id = uId.String()
	return
}
