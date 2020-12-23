package service

import (
	"banking/Domain"
	"banking/dto"
	"banking/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo Domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	err := request.Validate()
	if err != nil {
		return nil, err
	}

	acc := Domain.Account{
		AccountID:   "",
		CustomerID:  request.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "1",
	}

	newAccount, err := s.repo.Save(acc)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToResponseDTO()

	return &response, nil
}

func NewAccountService(repo Domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
