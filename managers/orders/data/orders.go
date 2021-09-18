package data

type Order struct {
	OrderNumber  int    `json:"orderNumber"`
	ItemId       int    `json:"itemId"`
	Price        int    `json:"price"`
	CustomerId   int    `json:"customerId"`
	CustomerName string `json:"customerName"`
	CustomerCity string `json:"customerCity"`
	VehicleId    int    `json:"vehicleId"`
	IsDelivered  bool   `json:"isDelivered"`
}
