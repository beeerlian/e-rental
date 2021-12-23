package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Motor struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type         string             `json:"type,omitempty" bson:"type,omitempty"`
	PoliceNumber string             `json:"police_number,omitempty" bson:"police_number,omitempty"`
	Merk         string             `json:"merk,omitempty" bson:"merk,omitempty"`
	Lecturer     string             `json:"lecturer,omitempty" bson:"lecturer,omitempty"`
	Participant  []MotorActivity    `json:"participant,omitempty" bson:"participant,omitempty"`
}
