package service

import (
	"banking/Domain"
	"banking/errs"
)

//user interface

type CustomerService interface {
	GetAllCustomers(string) ([]Domain.Customer, *errs.AppError)
	GetCustomerByID(string) (*Domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo Domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]Domain.Customer, *errs.AppError) {

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomerByID(id string) (*Domain.Customer, *errs.AppError) {
	return s.repo.ByID(id)
}

func NewCustomerService(repository Domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
