package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tutorial struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Uid         primitive.ObjectID `json:"uid,omitempty"`
	Lid         primitive.ObjectID `json:"lid,omitempty"`
	Title       string             `json:"title,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" validate:"required"`
	Price       int                `json:"pricw,omitempty" validate:"required"`
}
