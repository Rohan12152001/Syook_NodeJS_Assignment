package main

import (
	"github.com/Rohan12152001/Syook_Assignment/endpoints/auth"
	"github.com/Rohan12152001/Syook_Assignment/endpoints/customers"
	"github.com/Rohan12152001/Syook_Assignment/endpoints/items"
	"github.com/Rohan12152001/Syook_Assignment/endpoints/orders"
	"github.com/Rohan12152001/Syook_Assignment/endpoints/vehicles"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	server := gin.Default()

	auth := auth.New()
	auth.SetRoutes(server)

	items := items.New()
	items.SetRoutes(server)

	vehicles := vehicles.New()
	vehicles.SetRoutes(server)

	orders := orders.New()
	orders.SetRoutes(server)

	customers := customers.New()
	customers.SetRoutes(server)

	server.Run()
}
