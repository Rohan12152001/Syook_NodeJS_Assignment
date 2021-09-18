package customers

import (
	"context"
	"github.com/Rohan12152001/Syook_Assignment/managers/Customers/db"
	"github.com/Rohan12152001/Syook_Assignment/managers/customers/data"
	"github.com/Rohan12152001/Syook_Assignment/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type manager struct {
	db db.CustomersDBManager
}

func New() CustomersManager {
	return manager{
		db: db.New(),
	}
}

var logger = logrus.New()

func (m manager) CreateCustomer(ctx context.Context, customer data.Customer) (int, error) {
	// Call db
	customerId, err := m.db.CreateCustomer(customer)
	if err != nil {
		logger.Error("err: ", err)
		return -1, err
	}

	return customerId, nil
}

func (m manager) GetCustomer(ctx context.Context, Id int) (*data.Customer, error) {
	// Call db
	customer, err := m.db.GetCustomer(Id)
	if err != nil {
		logger.Error("err: ", err)
		if xerrors.Is(err, utils.NoRowsFound) {
			return nil, CustomerNotFound
		}
		return nil, err
	}
	return customer, nil
}
