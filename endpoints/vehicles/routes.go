package vehicles

import (
	"encoding/json"
	"fmt"
	"github.com/Rohan12152001/Syook_Assignment/managers/middleware"
	"github.com/Rohan12152001/Syook_Assignment/managers/vehicles"
	"github.com/Rohan12152001/Syook_Assignment/managers/vehicles/data"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strconv"
)

type Vehicles struct {
	vehicleManager vehicles.VehiclesManager
	authManager    middleware.AuthManager
}

func New() Vehicles {
	return Vehicles{
		vehicleManager: vehicles.New(),
		authManager:    middleware.New(),
	}
}

func getParam(c *gin.Context, paramName string) string {
	return c.Params.ByName(paramName)
}

var logger = logrus.New()

func (V Vehicles) SetRoutes(router *gin.Engine) {
	router.GET("/vehicles", V.authManager.AuthMiddleWareWithUser, V.GetAllVehicles)
	router.POST("/vehicle", V.authManager.AuthMiddleWareWithUser, V.CreateVehicle)
	router.GET("/vehicle/:id", V.authManager.AuthMiddleWareWithUser, V.ReadVehicle)
	router.PUT("/vehicle/:id", V.authManager.AuthMiddleWareWithUser, V.UpdateVehicle)
}

func (V Vehicles) GetAllVehicles(context *gin.Context) {
	// Call manager
	vehicles, err := V.vehicleManager.GetAllVehicles(context)
	if err != nil {
		// errors
		fmt.Println(err)
		context.AbortWithStatus(500)
	}

	context.JSON(200, gin.H{
		"vehicles": vehicles,
	})
}

func (V Vehicles) CreateVehicle(context *gin.Context) {
	vehiclePayload := data.VehicleStruct{}
	b, err := ioutil.ReadAll(context.Request.Body)

	err = json.Unmarshal(b, &vehiclePayload)
	if err != nil {
		context.AbortWithStatus(500)
		context.Error(err)
		return
	}

	// Manager
	vehicleId, err := V.vehicleManager.CreateVehicle(context, vehiclePayload)
	if err != nil {
		context.AbortWithStatus(500)
		context.Error(err)
		return
	}

	context.JSON(200, gin.H{
		"vehicleId": vehicleId,
	})
}

func (V Vehicles) ReadVehicle(context *gin.Context) {
	IdFromParam := getParam(context, "id")
	Id, err := strconv.Atoi(IdFromParam)
	if err != nil {
		// errors
		fmt.Println(err)
		context.AbortWithStatus(500)
	}

	// Manager
	vehicle, err := V.vehicleManager.GetVehicle(context, Id)
	if err != nil {
		// errors
		fmt.Println(err)
		context.AbortWithStatus(500)
	}

	if len(vehicle) == 0 {
		context.JSON(404, gin.H{
			"vehicle": "Not found",
		})
	} else {
		context.JSON(200, gin.H{
			"vehicle": vehicle[0],
		})
	}
}

func (V Vehicles) UpdateVehicle(context *gin.Context) {
	vehiclePayload := data.VehicleStruct{}
	b, err := ioutil.ReadAll(context.Request.Body)

	err = json.Unmarshal(b, &vehiclePayload)
	if err != nil {
		context.AbortWithStatus(500)
		context.Error(err)
		return
	}

	// Get ID from pathParam
	IdFromParam := getParam(context, "id")
	Id, err := strconv.Atoi(IdFromParam)
	if err != nil {
		// errors
		logger.Error("err: ", err)
		context.AbortWithStatus(500)
	}

	// Update struct with ItemId
	vehiclePayload.VehicleId = Id

	// Manager
	ok, err := V.vehicleManager.UpdateVehicle(context, vehiclePayload)
	if err != nil {
		// errors
		fmt.Println(err)
		context.AbortWithStatus(500)
	}

	if !ok {
		fmt.Println("Item not found!")
		context.JSON(404, gin.H{
			"vehicle": "Not found",
		})
		context.AbortWithStatus(404)
	} else {
		context.JSON(200, gin.H{
			"vehicle": "Updated",
		})
	}
}
