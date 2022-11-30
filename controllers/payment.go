package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nozlapp.com/modules/config"
	"nozlapp.com/modules/constants"
	"nozlapp.com/modules/db/actions"
	"nozlapp.com/modules/db/models"
	"nozlapp.com/modules/utils"
)

func CreateStripePaymentIntent(rs http.ResponseWriter, rq *http.Request) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var paymentIntent constants.PaymentIntent
	defer cancel()

	// initialize stripe key
	stripe.Key = config.GetEnv("STRIPE_KEY")

	//validate the request body
	if err := json.NewDecoder(rq.Body).Decode(&paymentIntent); err != nil {
		utils.ResponseEncoders(rs, http.StatusBadRequest, err.Error(), "error")
		return
	}

	//create payment intent object
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(paymentIntent.Amount),
		Currency: stripe.String(paymentIntent.Currency),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	}
	pi, err := paymentintent.New(params)

	if err != nil {
		utils.ResponseEncoders(rs, http.StatusInternalServerError, err.Error(), "error")
		return
	} else {
		utils.ResponseEncoders(rs, http.StatusOK, pi, "succuess")
	}
}

func SaveUserPaymentInfo(rs http.ResponseWriter, rq *http.Request) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var payment models.Payment
	defer cancel()

	//validate the request body
	if err := json.NewDecoder(rq.Body).Decode(&payment); err != nil {
		utils.ResponseEncoders(rs, http.StatusBadRequest, err.Error(), "error")
		return
	}

	newPayment := models.Payment{
		Id:                    primitive.NewObjectID(),
		PaidAmount:            payment.PaidAmount,
		StripeReferenceId:     payment.StripeReferenceId,
		Status:                payment.Status,
		Currency:              payment.Currency,
		CancellationReason:    payment.CancellationReason,
		StripePaymentMethodId: payment.StripePaymentMethodId,
		UserId:                payment.UserId,
		CreatedAt:             time.Now(),
	}

	inserted, err := actions.SavePaymentDetails(newPayment)

	if err != nil {
		utils.ResponseEncoders(rs, http.StatusInternalServerError, err.Error(), "error")
		return
	} else {
		utils.ResponseEncoders(rs, http.StatusCreated, inserted, "succuess")
		return
	}
}
