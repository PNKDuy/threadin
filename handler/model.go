package handler

type PaymentNotification struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Payer    string  `json:"payer"`
	Status   string  `json:"status"`
}
