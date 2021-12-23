package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RentDate   string             `json:"rent_date,omitempty" bson:"rent_date,omitempty"`
	ReturnDate string             `json:"return_date,omitempty" bson:"return_date,omitempty"`
	Customer   Customer           `json:"customer,omitempty" bson:"customer,omitempty"`
	Motor      Motor              `json:"motor,omitempty" bson:"motor,omitempty"`
}
