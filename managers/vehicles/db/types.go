package db

import (
	"context"
	"github.com/Rohan12152001/Syook_Assignment/managers/vehicles/data"
)

type VehiclesDBManager interface {
	GetAllVehicles() ([]data.VehicleStruct, error)
	GetVehicle(Id int) ([]data.VehicleStruct, error)
	CreateVehicle(newVehicle data.VehicleStruct) (int, error)
	UpdateVehicle(updatedVehicle data.VehicleStruct) error
	GetAvailableVehicleIds(matchCity string) (vehicles []data.VehicleStruct, err error) // [[Id, ordercount]]
	ReserveVehicleForOrderForCity(ctx context.Context, city string) (vehicleId int, err error)
	DecreaseActiveOrderCount(ctx context.Context, vehicleId int) (err error)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "#yourPassword"
	dbname   = "Logistics"
)
