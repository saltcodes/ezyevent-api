package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Organization struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	UserId      string    `json:"userId" gorm:"user_id"`
	BannerUrl   string    `json:"bannerUrl" gorm:"banner_url"`
	OrgImageUrl string    `json:"orgImageUrl" gorm:"org_image_url"`
	User        *User     `json:"user,omitempty" gorm:"foreignKey:UserId"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (base *Organization) BeforeCreate(db *gorm.DB) (err error) {
	uId, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	base.Id = uId.String()
	return
}
