package model

type User struct {
	Id         string `gorm:"primary_key" json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	ProfileUrl string `json:"profile_url"`
	DOB        int64  `json:"dob"`
}
