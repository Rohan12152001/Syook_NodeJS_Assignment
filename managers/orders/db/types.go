package db

import (
	"context"
	"github.com/Rohan12152001/Syook_Assignment/managers/orders/data"
)

type OrdersDBManager interface {
	StartTransactionInContext(ctx context.Context) (context.Context, error)
	CommitTransactionInContext(ctx context.Context) error
	RollbackTransactionInContext(ctx context.Context) error
	CreateOrder(ctx context.Context, order data.Order) (Id int, err error)
	UpdateVehicleForOrder(ctx context.Context, orderId, vehicleId int) (err error)
	UpdateIsDeliveredForOrder(ctx context.Context, orderId int) (err error)
	GetOrder(ctx context.Context, orderId int) (order *data.Order, err error)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "#yourPassword"
	dbname   = "Logistics"
)
