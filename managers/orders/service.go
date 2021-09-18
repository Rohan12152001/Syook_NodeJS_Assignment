package orders

import (
	"context"
	"github.com/Rohan12152001/Syook_Assignment/managers/Orders/db"
	"github.com/Rohan12152001/Syook_Assignment/managers/customers"
	customerData "github.com/Rohan12152001/Syook_Assignment/managers/customers/data"
	"github.com/Rohan12152001/Syook_Assignment/managers/items"
	"github.com/Rohan12152001/Syook_Assignment/managers/orders/data"
	"github.com/Rohan12152001/Syook_Assignment/managers/vehicles"
	"github.com/Rohan12152001/Syook_Assignment/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type manager struct {
	db              db.OrdersDBManager
	itemsManager    items.ItemsManager
	vehiclesManger  vehicles.VehiclesManager
	customersManger customers.CustomersManager
}

func New() OrdersManager {
	return manager{
		db:              db.New(),
		itemsManager:    items.New(),
		vehiclesManger:  vehicles.New(),
		customersManger: customers.New(),
	}
}

var logger = logrus.New()

func (m manager) CreateOrder(ctx context.Context, order data.Order) (OrderId int, err error) {
	orderId := -1

	// check if item exists & fetch price
	item, err := m.itemsManager.GetItem(ctx, order.ItemId)
	if err != nil {
		logger.Error("err: ", err)
		return orderId, err
	}

	// If yes then take the price, insert in the struct
	order.Price = item.Price

	// check if customerId exists (if orderId given in struct)
	var customerDetails *customerData.Customer
	if order.CustomerId != 0 { // since orderId will be zero if not filled
		customerDetails, err = m.customersManger.GetCustomer(ctx, order.CustomerId)
		if err != nil {
			logger.Error("err: ", err)
			return orderId, err
		}

	} else {
		// else create new customer & feed the ID
		customerDetails = &customerData.Customer{
			Name: order.CustomerName,
			City: order.CustomerCity}

		customerId, err := m.customersManger.CreateCustomer(ctx, *customerDetails)
		if err != nil {
			logger.Error("err: ", err)
			return orderId, err
		}
		order.CustomerId = customerId
	}

	// 1. Create order with vehicleId > Empty
	// 2. Assign vehicle
	// 3. orderUpdate
	// NOTE: 2 and 3 in transaction

	// CreateOrder so we dont miss any orders
	orderId, err = m.db.CreateOrder(ctx, order)
	if err != nil {
		logger.Error("err: ", err)
		return orderId, err
	}

	// Using transactions in on service layer (since we must take care of Vehicles & Orders Table together)
	ctx, err = m.db.StartTransactionInContext(ctx)
	if err != nil {
		logger.Error("err: ", err)
		return orderId, err
	}

	// In case of failure Rollback
	defer func() {
		if err != nil {
			logger.Error("err: ", err)
			m.db.RollbackTransactionInContext(ctx)
		}
	}()

	vehicleId, err := m.vehiclesManger.ReserveVehicleForOrderForCity(ctx, order.CustomerCity)
	if err != nil {
		logger.Error("err: ", err)
		return orderId, err
	}
	order.VehicleId = vehicleId

	err = m.db.UpdateVehicleForOrder(ctx, orderId, vehicleId)
	if err != nil {
		logger.Error("err: ", err)
		return orderId, err
	}

	err = m.db.CommitTransactionInContext(ctx)
	if err != nil {
		logger.Error("err: ", err)
		return orderId, err
	}

	return orderId, nil
}

func (m manager) OrderDelivered(ctx context.Context, orderId int) (err error) {
	// get order
	order, err := m.db.GetOrder(ctx, orderId)
	if err != nil {
		logger.Error("err: ", err)
		return err
	}
	if order.VehicleId == 0 {
		return NoVehicleAssigned
	}
	if order.IsDelivered {
		return nil
	}

	// Using transactions in on service layer (take care of Orders & vehicles both tables)
	ctx, err = m.db.StartTransactionInContext(ctx)
	if err != nil {
		logger.Error("err: ", err)
		return err
	}

	// In case of failure Rollback
	defer func() {
		if err != nil {
			logger.Error("err: ", err)
			m.db.RollbackTransactionInContext(ctx)
		}
	}()

	// First change isDelivered
	err = m.db.UpdateIsDeliveredForOrder(ctx, orderId)
	if err != nil {
		logger.Error("err: ", err)
		// Order not found (sql.ErrNoRows) OR already delivered
		if xerrors.Is(err, utils.NorowsUpdated) {
			return CannotUpdateIsDelivered
		}
		return err // Or Maybe Already delivered
	}

	// Then change activeCount
	err = m.vehiclesManger.DecreaseActiveOrderCount(ctx, order.VehicleId)
	if err != nil {
		logger.Error("err: ", err)
		return err
	}

	// commit
	err = m.db.CommitTransactionInContext(ctx)
	if err != nil {
		logger.Error("err: ", err)
		return err
	}

	return nil
}
