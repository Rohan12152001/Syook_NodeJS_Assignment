package items

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/Rohan12152001/Syook_Assignment/managers/items"
	"github.com/Rohan12152001/Syook_Assignment/managers/items/data"
	"github.com/Rohan12152001/Syook_Assignment/managers/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"io/ioutil"
	_ "io/ioutil"
	"strconv"
)

type Items struct {
	itemManager items.ItemsManager
	authManager middleware.AuthManager
}

var logger = logrus.New()

func New() Items {
	return Items{
		itemManager: items.New(),
		authManager: middleware.New(),
	}
}

func getParam(c *gin.Context, paramName string) string {
	return c.Params.ByName(paramName)
}

func (I Items) SetRoutes(router *gin.Engine) {
	router.GET("/items", I.authManager.AuthMiddleWareWithUser, I.GetAllItems)
	router.POST("/item", I.authManager.AuthMiddleWareWithUser, I.CreateItem)
	router.GET("/item/:id", I.authManager.AuthMiddleWareWithUser, I.ReadItem)
	router.PUT("/item/:id", I.authManager.AuthMiddleWareWithUser, I.UpdateItem)
}

// Handlers
func (I Items) GetAllItems(context *gin.Context) {
	// Call manager
	items, err := I.itemManager.GetAllItems(context)
	if err != nil {
		// errors
		context.AbortWithStatus(500)
	}

	context.JSON(200, gin.H{
		"items": items,
	})

}

func (I Items) ReadItem(context *gin.Context) {
	IdFromParam := getParam(context, "id")
	Id, err := strconv.Atoi(IdFromParam)
	if err != nil {
		// errors
		logrus.Error("err: ", err)
		context.AbortWithStatus(500)
	}

	// Manager
	item, err := I.itemManager.GetItem(context, Id)
	if err != nil {
		// errors
		if xerrors.Is(err, items.ItemNotFound) {
			context.JSON(404, gin.H{
				"item": "Not found",
			})
			return
		}
		context.AbortWithStatus(500)
	}

	context.JSON(200, gin.H{
		"item": item,
	})
}

func (I Items) CreateItem(context *gin.Context) {
	itemPayload := data.Item{}
	b, err := ioutil.ReadAll(context.Request.Body)

	err = json.Unmarshal(b, &itemPayload)
	if err != nil {
		logger.Error("err: ", err)
		context.AbortWithStatus(500)
		return
	}

	// Manager
	itemId, err := I.itemManager.CreateItem(context, itemPayload.Name, itemPayload.Price)
	if err != nil {
		context.AbortWithStatus(500)
		context.Error(err)
		return
	}

	context.JSON(200, gin.H{
		"itemId": itemId,
	})

}

func (I Items) UpdateItem(context *gin.Context) {
	itemPayload := data.Item{}
	b, err := ioutil.ReadAll(context.Request.Body)

	err = json.Unmarshal(b, &itemPayload)
	if err != nil {
		logger.Error("err: ", err)
		context.AbortWithStatus(500)
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
	itemPayload.ItemId = Id

	// Manager
	ok, err := I.itemManager.UpdateItem(context, itemPayload)
	if err != nil {
		// errors
		context.AbortWithStatus(500)
	}

	if !ok {
		fmt.Println("Item not found!")
		context.JSON(404, gin.H{
			"item": "Not found",
		})
		context.AbortWithStatus(404)
	} else {
		context.JSON(200, gin.H{
			"item": "Updated",
		})
	}
}
