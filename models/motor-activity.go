package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MotorActivity struct {
	CustomerId primitive.ObjectID `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	Email      string             `json:"email,omitempty" bson:"email,omitempty"`
	Attende    string             `json:"attende,omitempty" bson:"attende,omitempty"`
}
