package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Link struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Option string             `json:"option,omitempty" validate:"required"`
}
