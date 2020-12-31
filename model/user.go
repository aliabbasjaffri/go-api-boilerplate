package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string        		`json:"name" bson:"name"`
	Age   int           		`json:"age" bson:"age"`
	Email string        		`json:"email" bson:"email"`
}