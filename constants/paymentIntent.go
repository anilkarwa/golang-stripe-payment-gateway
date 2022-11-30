package constants

type PaymentIntent struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}
