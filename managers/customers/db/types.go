package db

import (
	"github.com/Rohan12152001/Syook_Assignment/managers/customers/data"
)

type CustomersDBManager interface {
	CreateCustomer(customer data.Customer) (int, error)
	GetCustomer(Id int) (*data.Customer, error)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "#yourPassword"
	dbname   = "Logistics"
)
