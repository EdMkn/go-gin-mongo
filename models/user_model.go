package models

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Friend struct {
	Id   int64  `json:"id" bson:"id,omitempty"`
	Name string `json:"name" bson:"name,omitempty"`
}

type User struct {
	Id         string   `json:"id,omitempty" bson:"id,omitempty"`
	Password   string   `json:"password" bson:"password" validate:"required"`
	IsActive   *bool    `json:"isActive" bson:"isActive,omitempty" validate:"required"`
	Balance    string   `json:"balance" bson:"balance" validate:"required"`
	Age        uint8    `json:"age,omitempty" bson:"age,omitempty" validate:"required,gte=0,lte=130"`
	Name       string   `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Gender     Gender   `json:"gender,omitempty" bson:"gender,omitempty" validate:"required"`
	Company    string   `json:"company,omitempty" bson:"company,omitempty" validate:"required"`
	Email      string   `json:"email,omitempty" bson:"email,omitempty" validate:"email"`
	Phone      string   `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	Address    string   `json:"address,omitempty" bson:"address,omitempty" validate:"required"`
	About      string   `json:"about,omitempty" bson:"about,omitempty" `
	Registered string   `json:"registered,omitempty" bson:"registered,omitempty" `
	Latitude   float32  `json:"latitude,omitempty" bson:"latitude,omitempty" validate:"numeric" `
	Longitude  float32  `json:"longitude,omitempty" bson:"longitude,omitempty" validate:"numeric"`
	Tags       []string `json:"tags,omitempty" bson:"tags,omitempty" `
	Friends    []Friend `json:"friends" bson:"friends,omitempty" `
	Data       string   `json:"data,omitempty" bson:"data,omitempty" validate:"required"`
}
