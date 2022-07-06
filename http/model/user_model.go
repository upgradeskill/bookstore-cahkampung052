package model

type User struct {
	Id    int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"size:150;unique:true" json:"name"`
	Email string `gorm:"size:150;unique:true" json:"email"`
	Roles string `json:"roles"`
}

type UserPayload struct {
	Id    int         `json:"id"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Roles interface{} `json:"roles"`
}
