package model

type Auth struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
