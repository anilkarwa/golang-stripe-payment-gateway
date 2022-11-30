package actions

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"nozlapp.com/modules/db"
	"nozlapp.com/modules/db/models"
)

var paymentCollection *mongo.Collection = db.GetCollection(db.MongoClient, "payments")

func SavePaymentDetails(payment models.Payment) (any, error) {
	inserted, err := paymentCollection.InsertOne(context.Background(), payment)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}
