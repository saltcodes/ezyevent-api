package model

type Comment struct {
	Id        string `gorm:"PrimaryKey" json:"id"`
	Message   string `json:"message"`
	EventId   string `json:"event_id"`
	UserId    string `json:"userId" gorm:"user_id"`
	User      *User  `json:"user" gorm:"foreignKey:UserId"`
	UpdatedAt int64  `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
	CreatedAt int64  `gorm:"autoCreateTime:nano"`
}
