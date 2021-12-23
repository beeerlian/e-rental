package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerActivity struct {
	MotorId   primitive.ObjectID `json:"motor_id,omitempty" bson:"motor_id,omitempty"`
	MotorType string             `json:"motor_type,omitempty" bson:"motor_type,omitempty"`
	Attende   string             `json:"attende,omitempty" bson:"attende,omitempty"`
}
