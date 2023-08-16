package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Restaurants struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Borough string             `json:"borough" bson:"borough,required"`
	Cuisine string             `json:"cuisine" bson:"cuisine,required"`
}
