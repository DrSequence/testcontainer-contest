package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Portfolio struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name    string             `bson:"name" json:"name"`
	Details string             `bson:"details" json:"details"`
}
