package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Fullname string             `json:"fullname,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required,email"`
	Password string             `json:"password,omitempty" validate:"required,min=4,max=15"`
}
