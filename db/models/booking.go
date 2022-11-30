package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Booking struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id"`
	OperatorUserId User               `json:"operatorUserId"`
	Rate           float64            `json:"rate"`
	FboId          Fbo                `json:"fboId"`
	OperatorId     Operator           `json:"operatorId"`
}
