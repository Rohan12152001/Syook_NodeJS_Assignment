package db

import (
	_ "context"
	"fmt"
	"github.com/Rohan12152001/Syook_Assignment/managers/customers/data"
	"github.com/Rohan12152001/Syook_Assignment/utils"
	"github.com/jmoiron/sqlx"
)

type manager struct {
	db *sqlx.DB
}

func New() CustomersDBManager {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return manager{
		db: db,
	}
}

func (m manager) CreateCustomer(customer data.Customer) (int, error) {
	fmt.Println(">>>> cutomer", customer)
	query := "Insert into customers (name, city) values($1, $2) RETURNING customerId;"

	Id := -1
	err := m.db.QueryRow(query,
		customer.Name,
		customer.City).Scan(&Id)
	if err != nil {
		return Id, err
	}

	return Id, nil
}

func (m manager) GetCustomer(Id int) (*data.Customer, error) {
	query := "Select * from customers where customerId=$1"
	customers := []*data.Customer{}

	err := m.db.Select(&customers, query, Id)
	if err != nil {
		return nil, err
	}

	if len(customers) == 0 {
		return nil, utils.NoRowsFound
	}
	return customers[0], nil
}
