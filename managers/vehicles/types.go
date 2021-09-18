package vehicles

import (
	"context"
	"fmt"
	"github.com/Rohan12152001/Syook_Assignment/managers/vehicles/data"
)

var (
	NoVehicleFound      = fmt.Errorf("no vehicles found")
	CannotDecreaseCount = fmt.Errorf("cannot decrease count!")
)

type VehiclesManager interface {
	GetAllVehicles(ctx context.Context) ([]data.VehicleStruct, error)
	GetVehicle(ctx context.Context, Id int) ([]data.VehicleStruct, error)
	CreateVehicle(ctx context.Context, newVehicle data.VehicleStruct) (int, error)
	UpdateVehicle(ctx context.Context, updatedVehicle data.VehicleStruct) (bool, error)
	ReserveVehicleForOrderForCity(ctx context.Context, city string) (vehicleId int, err error)
	DecreaseActiveOrderCount(ctx context.Context, vehicleId int) (err error)
}
