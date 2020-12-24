package Domain

import (
	"banking/dto"
	"banking/errs"
)

type Transaction struct {
	TransactionID   string
	AccountID       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

func (t Transaction) ToResponseDTO(balance float64) dto.TransactionResponse {
	return dto.TransactionResponse{Balance: balance, TransactionID: t.TransactionID}
}

type TransactionRepository interface {
	Record(Transaction) (*Transaction, *errs.AppError)
	GetBalance(string) (float64, *errs.AppError)
	UpdateBalance(Transaction) (float64, *errs.AppError)
	Rollback(transaction Transaction) *errs.AppError
}
