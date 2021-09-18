package customers

import (
	"encoding/json"
	"github.com/Rohan12152001/Syook_Assignment/managers/customers"
	"github.com/Rohan12152001/Syook_Assignment/managers/customers/data"
	"github.com/Rohan12152001/Syook_Assignment/managers/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"io/ioutil"
	"strconv"
)

type Customers struct {
	customerManager customers.CustomersManager
	authManager     middleware.AuthManager
}

func New() Customers {
	return Customers{
		customerManager: customers.New(),
		authManager:     middleware.New(),
	}
}

func getParam(c *gin.Context, paramName string) string {
	return c.Params.ByName(paramName)
}

func (C Customers) SetRoutes(router *gin.Engine) {
	router.POST("/customer", C.authManager.AuthMiddleWareWithUser, C.CreateCustomer)
	router.GET("/customer/:id", C.authManager.AuthMiddleWareWithUser, C.GetCustomer)
}

var logger = logrus.New()

func (C Customers) CreateCustomer(context *gin.Context) {
	customerPayload := data.Customer{}
	b, err := ioutil.ReadAll(context.Request.Body)

	err = json.Unmarshal(b, &customerPayload)
	if err != nil {
		logger.Error("err: ", err)
		context.AbortWithStatus(500)
		return
	}

	// Manager
	customerId, err := C.customerManager.CreateCustomer(context, customerPayload)
	if err != nil {
		context.AbortWithStatus(500)
		context.Error(err)
		return
	}

	context.JSON(200, gin.H{
		"customerId": customerId,
	})
}

func (C Customers) GetCustomer(context *gin.Context) {
	IdFromParam := getParam(context, "id")
	Id, err := strconv.Atoi(IdFromParam)
	if err != nil {
		// errors
		logger.Error("err: ", err)
		context.AbortWithStatus(500)
	}

	// Manager
	customer, err := C.customerManager.GetCustomer(context, Id)
	if err != nil {
		// errors
		if xerrors.Is(err, customers.CustomerNotFound) {
			context.JSON(404, gin.H{
				"customer": "Not found",
			})
			return
		}
		context.AbortWithStatus(500)
	}

	context.JSON(200, gin.H{
		"customer": customer,
	})

}
