package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Motor struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type         string             `json:"type,omitempty" bson:"type,omitempty"`
	PoliceNumber string             `json:"pnumber,omitempty" bson:"pnumber,omitempty"`
	Merk         string             `json:"merk,omitempty" bson:"merk,omitempty"`
	Lecturer     string             `json:"lecturer,omitempty" bson:"lecturer,omitempty"`
}
