package model

type Comment struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Message   string `json:"message"`
	EventId   string `json:"event_id"`
	User      User   `gorm:"foreignkey:Id;references:user_id" json:"user"`
	UpdatedAt int64  `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
	CreatedAt int64  `gorm:"autoCreateTime:nano"`
}
