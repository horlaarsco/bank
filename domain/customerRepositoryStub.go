package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "John", "New York", 10001, "01/01/1990", true},
		{"2", "Jane", "New York", 10002, "01/01/1990", true},
		{"3", "Romi", "New York", 10003, "01/01/1990", true},
	}

	return CustomerRepositoryStub{customers}
}
