package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	ID         string    `gorm:"id" json:"ID"`
	ScheduleID string    `gorm:"schedule_id" json:"scheduleId"`
	UserID     string    `gorm:"user_id" json:"userId"`
	User       *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Schedule   *Schedule `json:"schedule,omitempty" gorm:"foreignKey:ScheduleID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func (base *Booking) BeforeCreate(db *gorm.DB) (err error) {
	uId, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	base.ID = uId.String()
	return
}
