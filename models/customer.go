package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone         string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Password      string             `json:"password,omitempty" bson:"password,omitempty"`
	Status        string             `json:"status,omitempty" bson:"status,omitempty"`
	Activities    []CustomerActivity `json:"activities,omitempty" bson:"activities,omitempty"`
	MotorsCreated []Motor            `json:"motors_created,omitempty" bson:"motors_created,omitempty"`
}
