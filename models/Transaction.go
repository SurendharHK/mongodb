package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct{

	ID primitive.ObjectID `json:"id" bson:"_id"`
	Account_Id int `json:"account_id" bson:"account_id"`
	Transaction_count int `json:"transaction_count" bson:"transaction_count"`



}
