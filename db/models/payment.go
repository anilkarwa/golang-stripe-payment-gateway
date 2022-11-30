package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	Id                    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PaidAmount            float64            `json:"paidAmount,omitempty"`
	StripeReferenceId     string             `json:"stripeReferenceId,omitempty"`
	Status                string             `json:"status,omitempty"`
	Currency              string             `json:"currency,omitempty"`
	CancellationReason    string             `json:"cancellationReason,omitempty"`
	StripePaymentMethodId string             `json:"stripePaymentMethodId,omitempty"`
	CreatedAt             time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
	UserId                primitive.ObjectID `json:"userId" bson:"userId"`
}
