package dto

type TransactionResponse struct {
	Balance       float64 `json:"balance"`
	TransactionID string  `json:"transaction_id"`
}
