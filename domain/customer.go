package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     int
	DateofBirth string
	Active      bool
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
