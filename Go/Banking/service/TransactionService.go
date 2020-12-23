package service

import (
	"banking/Domain"
	"banking/dto"
	"banking/errs"
	"strings"
	"time"
)

type TransactionService interface {
	NewTransaction(dto.NewTransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo Domain.TransactionRepository
}

func (t DefaultTransactionService) NewTransaction(request dto.NewTransactionRequest) (*dto.TransactionResponse, *errs.AppError) {

	balance, err := t.repo.GetBalance(request.AccountID)
	if err != nil {
		return nil, err
	}
	err = request.Validate(balance)
	if err != nil {
		return nil, err
	}

	transaction := Domain.Transaction{
		TransactionID:   "",
		AccountID:       request.AccountID,
		Amount:          request.Amount,
		TransactionType: request.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	if strings.ToLower(transaction.TransactionType) == "withdrawal" {
		transaction.Amount = -transaction.Amount
	}

	newTransaction, err := t.repo.Record(transaction)
	if err != nil {
		return nil, err
	}

	balance, err = t.repo.UpdateBalance(transaction)
	if err != nil {
		return nil, err
	}

	response := newTransaction.ToResponseDTO(balance)

	return &response, nil
}

func NewTransactionService(repo Domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo}
}
