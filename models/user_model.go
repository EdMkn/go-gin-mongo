package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type User struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	Password   string             `json:"password,omitempty" validate:"required"`
	IsActive   bool               `json:"isActive,omitempty" validate:"required"`
	Balance    float32            `json:"balance,omitempty" validate:"required"`
	Age        int                `json:"age,omitempty" validate:"required"`
	Name       string             `json:"name,omitempty" validate:"required"`
	Gender     Gender             `json:"gender,omitempty" validate:"required"`
	Company    string             `json:"company,omitempty" validate:"required"`
	Email      string             `json:"email,omitempty" validate:"required"`
	Phone      string             `json:"phone,omitempty" validate:"required"`
	Address    string             `json:"address,omitempty" validate:"required"`
	About      string             `json:"about,omitempty" validate:"required"`
	Registered string             `json:"registered,omitempty" validate:"required"`
	Latitude   float32            `json:"latitude,omitempty" validate:"required"`
	Longitude  float32            `json:"longitude,omitempty" validate:"required"`
	Tags       []string           `json:"tags,omitempty" validate:"required"`
	Friends    []string           `json:"friends"`
	Data       string             `json:"data,omitempty" validate:"required"`
}
