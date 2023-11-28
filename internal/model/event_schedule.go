package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventSchedule struct {
	Id        string `json:"id"`
	EventId   string `gorm:"event_id"`
	Date      string `json:"date"`
	OrgId     string `gorm:"org_id"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli"` // Use unix nano seconds as updating time
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
}

func (base *EventSchedule) BeforeCreate(db *gorm.DB) (err error) {
	uId, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	base.Id = uId.String()
	return
}
