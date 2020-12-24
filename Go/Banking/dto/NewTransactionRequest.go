package dto

import (
	"banking/errs"
	"strings"
)

type NewTransactionRequest struct {
	AccountID       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

func (r NewTransactionRequest) Validate(balance float64) *errs.AppError {
	if strings.ToLower(r.TransactionType) != "withdrawal" && strings.ToLower(r.TransactionType) != "deposit" {
		return errs.NewValidationError("transaction types should be withdrawal or deposit")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("amount must be positive")
	}
	if strings.ToLower(r.TransactionType) == "withdrawal" && (balance-r.Amount) < 0 {
		return errs.NewValidationError("insufficient funds on account")
	}
	return nil
}
