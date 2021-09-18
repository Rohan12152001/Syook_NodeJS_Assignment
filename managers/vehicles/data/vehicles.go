package data

type VehicleStruct struct {
	VehicleId          int    `json:"vehicleId"`
	RegistrationNumber string `json:"registrationNumber"`
	VehicleType        string `json:"vehicleType"`
	City               string `json:"city"`
	ActiveOrdersCount  int    `json:"activeOrdersCount"`
}
