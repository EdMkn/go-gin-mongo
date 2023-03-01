package models

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type User struct {
	Id         string   `json:"id,omitempty" bson:"id,omitempty"`
	Password   string   `json:"password,omitempty" bson:"password,omitempty" validate:"required"`
	IsActive   *bool    `json:"isActive,omitempty" bson:"isActive,omitempty" validate:"required"`
	Balance    string   `json:"balance,omitempty" bson:"balance,omitempty" validate:"required"`
	Age        int      `json:"age,omitempty" bson:"age,omitempty" validate:"required"`
	Name       string   `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Gender     Gender   `json:"gender,omitempty" bson:"gender,omitempty" validate:"required"`
	Company    string   `json:"company,omitempty" bson:"company,omitempty" validate:"required"`
	Email      string   `json:"email,omitempty" bson:"email,omitempty" validate:"required"`
	Phone      string   `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	Address    string   `json:"address,omitempty" bson:"address,omitempty" validate:"required"`
	About      string   `json:"about,omitempty" bson:"about,omitempty" validate:"required"`
	Registered string   `json:"registered,omitempty" bson:"registered,omitempty" `
	Latitude   float32  `json:"latitude,omitempty" bson:"latitude,omitempty" `
	Longitude  float32  `json:"longitude,omitempty" bson:"longitude,omitempty" `
	Tags       []string `json:"tags,omitempty" bson:"tags,omitempty" `
	Friends    []string `json:"friends" bson:"friends,omitempty" `
	Data       string   `json:"data,omitempty" bson:"data,omitempty" validate:"required"`
}
