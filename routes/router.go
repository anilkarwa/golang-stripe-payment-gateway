package routes

import (
	"github.com/gorilla/mux"
	"nozlapp.com/modules/controllers"
)

func Router(router *mux.Router) {

	router.HandleFunc("/api/v1/payment-intent", controllers.CreateStripePaymentIntent).Methods("POST")
	router.HandleFunc("/api/v1/save-payment", controllers.SaveUserPaymentInfo).Methods("POST")
}
