package model

type Auth struct {
	UserId   string `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}
