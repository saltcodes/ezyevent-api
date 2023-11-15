package model

type Comment struct {
	Id        string `gorm:"PrimaryKey" json:"id"`
	UserId    string `json:"user_id" gorm:"foreignKey:user_id,"`
	Message   string `json:"message"`
	EventId   string `json:"event_id"`
	User      User   `json:"user"`
	UpdatedAt int64  `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
	CreatedAt int64  `gorm:"autoCreateTime:nano"`
}
