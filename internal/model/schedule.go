package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type EventSchedule struct {
	Id        string    `json:"id"`
	EventId   string    `gorm:"event_id"`
	Date      string    `json:"date"`
	OrgId     string    `gorm:"org_id"`
	Event     *Event    ` json:"event,omitempty" gorm:"event,foreignKey:EventId"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (base *EventSchedule) BeforeCreate(db *gorm.DB) (err error) {
	uId, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	base.Id = uId.String()
	return
}
