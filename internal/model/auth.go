package model

type Auth struct {
	UserId   string `bson:"_id,omitempty" json:"id"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
