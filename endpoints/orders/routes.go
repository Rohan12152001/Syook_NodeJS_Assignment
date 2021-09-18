package orders

import (
	"database/sql"
	"encoding/json"
	"github.com/Rohan12152001/Syook_Assignment/managers/customers"
	"github.com/Rohan12152001/Syook_Assignment/managers/items"
	"github.com/Rohan12152001/Syook_Assignment/managers/middleware"
	"github.com/Rohan12152001/Syook_Assignment/managers/orders"
	"github.com/Rohan12152001/Syook_Assignment/managers/orders/data"
	"github.com/Rohan12152001/Syook_Assignment/managers/vehicles"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"io/ioutil"
	"strconv"
)

type Orders struct {
	orderManager orders.OrdersManager
	authManager  middleware.AuthManager
}

func New() Orders {
	return Orders{
		orderManager: orders.New(),
		authManager:  middleware.New(),
	}
}

func getParam(c *gin.Context, paramName string) string {
	return c.Params.ByName(paramName)
}

var logger = logrus.New()

func (O Orders) SetRoutes(router *gin.Engine) {
	router.POST("/order", O.authManager.AuthMiddleWareWithUser, O.CreateOrder)
	router.POST("/order/delivered/:orderId", O.authManager.AuthMiddleWareWithUser, O.OrderDelivered) // Order Delivered API
}

func (O Orders) CreateOrder(context *gin.Context) {
	orderPayload := data.Order{}
	b, err := ioutil.ReadAll(context.Request.Body)

	err = json.Unmarshal(b, &orderPayload)
	if err != nil {
		context.AbortWithStatus(500)
		context.Error(err)
		return
	}

	// Manager
	orderId, err := O.orderManager.CreateOrder(context, orderPayload)
	if err != nil {
		if xerrors.Is(err, vehicles.NoVehicleFound) {
			context.JSON(404, gin.H{
				"error_message": "No vehicles found for order!",
			})
			return
		} else if xerrors.Is(err, customers.CustomerNotFound) {
			context.JSON(404, gin.H{
				"error_message": "Customer not found!",
			})
			return
		} else if xerrors.Is(err, items.ItemNotFound) {
			context.JSON(404, gin.H{
				"error_message": "Item not found!",
			})
			return
		} else if xerrors.Is(err, orders.CannotUpdateIsDelivered) {
			context.JSON(500, gin.H{
				"error_message": "Cannot Update!",
			})
			return
		} else if xerrors.Is(err, orders.NoVehicleAssigned) {
			context.JSON(400, gin.H{
				"error_message": "No vehicle assigned!",
			})
			return
		}
		context.AbortWithStatus(500)
		context.Error(err)
		return
	}

	context.JSON(200, gin.H{
		"orderId": orderId,
	})
}

func (O Orders) OrderDelivered(context *gin.Context) {
	// Path param
	IdFromParam := getParam(context, "orderId")
	OrderId, err := strconv.Atoi(IdFromParam)
	if err != nil {
		// errors
		logrus.Error("err: ", err)
		context.AbortWithStatus(500)
	}

	// Manager
	err = O.orderManager.OrderDelivered(context, OrderId)
	if err != nil {
		if xerrors.Is(err, sql.ErrNoRows) {
			context.JSON(404, gin.H{
				"error_message": "Order not found!",
			})
			return
		} else if xerrors.Is(err, vehicles.NoVehicleFound) {
			context.JSON(404, gin.H{
				"error_message": "Vehicle not found!",
			})
			return
		} else if xerrors.Is(err, vehicles.CannotDecreaseCount) {
			context.JSON(500, gin.H{
				"error_message": "Cannot decrease count!",
			})
			return
		}
		context.Error(err)
		context.AbortWithStatus(500)
	}

	context.JSON(200, gin.H{
		"Order": "Updated",
	})

}
