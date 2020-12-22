package service

import (
	"banking/Domain"
	"banking/dto"
	"banking/errs"
)

//user interface

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerByID(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo Domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDTO())
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomerByID(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repo.ByID(id)
	if err != nil {
		return nil, err
	}

	response := customer.ToDTO()

	return &response, nil
}

func NewCustomerService(repository Domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
