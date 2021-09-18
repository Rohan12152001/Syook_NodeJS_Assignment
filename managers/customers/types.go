package customers

import (
	"context"
	"fmt"
	"github.com/Rohan12152001/Syook_Assignment/managers/customers/data"
)

var (
	CustomerNotFound = fmt.Errorf("customer not found")
)

type CustomersManager interface {
	CreateCustomer(ctx context.Context, customer data.Customer) (int, error)
	GetCustomer(ctx context.Context, Id int) (*data.Customer, error)
}
