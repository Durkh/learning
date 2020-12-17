package Domain

//stub repository connected to the interface

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"0001", "Egídio", "Patos", "58000", "12/01", "1"},
		{"0002", "Dio", "Patos", "58000", "12/01", "1"},
	}

	return CustomerRepositoryStub{customers}
}
