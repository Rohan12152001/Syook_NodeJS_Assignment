package db

import (
	"context"
	_ "context"
	"database/sql"
	"errors"
	"fmt"
	//"github.com/Rohan12152001/Syook_Assignment/endpoints/orders"
	"github.com/Rohan12152001/Syook_Assignment/managers/orders/data"
	"github.com/Rohan12152001/Syook_Assignment/utils"
	"github.com/jmoiron/sqlx"
)

type manager struct {
	db *sqlx.DB
}

func New() OrdersDBManager {
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

func (m manager) CreateOrder(ctx context.Context, order data.Order) (OrderId int, err error) {
	query := "Insert into orders (itemId, price, customerId) values($1, $2, $3) RETURNING orderNumber;"

	Id := -1
	err = m.db.QueryRow(query,
		order.ItemId,
		order.Price,
		order.CustomerId).Scan(&Id)
	if err != nil {
		return Id, err
	}

	return Id, nil
}

func (m manager) UpdateVehicleForOrder(ctx context.Context, orderId, vehicleId int) (err error) {
	// 2. Order table vehicle id insert

	txn, ok := utils.GetTransactionFromContext(ctx)
	if !ok {
		txn, err = m.db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				txn.Rollback()
			} else {
				txn.Commit()
			}
		}()
	}

	query := "update orders set vehicleId = $1 where orderNumber = $2;"
	r, err := txn.Exec(query, vehicleId, orderId)
	if err != nil {
		return err
	}
	rowsAffected, err := r.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return utils.NorowsUpdated
	}

	return nil
}

func (m manager) CommitTransactionInContext(ctx context.Context) error {
	txn, ok := utils.GetTransactionFromContext(ctx)
	if !ok {
		return errors.New("transaction not found")
	}
	return txn.Commit()
}

func (m manager) RollbackTransactionInContext(ctx context.Context) error {
	txn, ok := utils.GetTransactionFromContext(ctx)
	if !ok {
		return errors.New("transaction not found")
	}
	return txn.Rollback()
}

func (m manager) StartTransactionInContext(ctx context.Context) (context.Context, error) {
	txn, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	ctx = utils.SetTransactionInContext(ctx, txn)
	return ctx, nil
}

func (m manager) UpdateIsDeliveredForOrder(ctx context.Context, orderId int) (err error) {

	txn, ok := utils.GetTransactionFromContext(ctx)
	if !ok {
		txn, err = m.db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				txn.Rollback()
			} else {
				txn.Commit()
			}
		}()
	}

	query := "update orders set isDelivered = true where orderNumber = $1 and vehicleId is NOT NULL and isDelivered=false;"
	var result sql.Result
	result, err = txn.Exec(query, orderId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return utils.NorowsUpdated
	}

	return nil
}

func (m manager) GetOrder(ctx context.Context, orderId int) (order *data.Order, err error) {
	query := "Select * from orders where orderNumber=$1"
	orders := []*data.Order{}

	err = m.db.Select(&orders, query, orderId)
	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, utils.NoRowsFound
	}

	return orders[0], nil
}
