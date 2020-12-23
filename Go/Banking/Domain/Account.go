package Domain

import (
	"banking/dto"
	"banking/errs"
)

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountID}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
