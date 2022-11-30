package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Email     string             `json:"email,omitempty"`
	UserType  int                `json:"userType"`
	Status    int                `json:"status"`
	IsDeleted bool               `json:"isDeleted"`
}
