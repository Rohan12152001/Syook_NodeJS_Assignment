package orders

import (
	"context"
	"fmt"
	"github.com/Rohan12152001/Syook_Assignment/managers/orders/data"
)

var (
	CannotUpdateIsDelivered = fmt.Errorf("cannot update!")
	NoVehicleAssigned       = fmt.Errorf("no vehicle assigned!")
)

type OrdersManager interface {
	CreateOrder(ctx context.Context, order data.Order) (OrderId int, err error)
	OrderDelivered(ctx context.Context, orderId int) (err error)
}
