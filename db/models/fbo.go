package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Fbo struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	FboName string             `json:"fboName"`
	UserId  User               `json:"userId"`
}
