package vehicles

import (
	"context"
	"github.com/Rohan12152001/Syook_Assignment/managers/vehicles/data"
	"github.com/Rohan12152001/Syook_Assignment/managers/vehicles/db"
	"github.com/Rohan12152001/Syook_Assignment/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type manager struct {
	db db.VehiclesDBManager
}

func New() VehiclesManager {
	return manager{
		db: db.New(),
	}
}

var logger = logrus.New()

func (m manager) GetAllVehicles(ctx context.Context) ([]data.VehicleStruct, error) {
	// Call db
	vehicles, err := m.db.GetAllVehicles()
	if err != nil {
		logger.Error("err: ", err)
		return nil, err
	}

	return vehicles, nil
}

func (m manager) GetVehicle(ctx context.Context, Id int) ([]data.VehicleStruct, error) {
	// Call db
	vehicle, err := m.db.GetVehicle(Id)
	if err != nil {
		logger.Error("err: ", err)
		return []data.VehicleStruct{}, err
	}

	return vehicle, nil
}

func (m manager) CreateVehicle(ctx context.Context, newVehicle data.VehicleStruct) (int, error) {
	// Call db
	vehId, err := m.db.CreateVehicle(newVehicle)
	if err != nil {
		logger.Error("err: ", err)
		return -1, err
	}

	return vehId, nil
}

func (m manager) UpdateVehicle(ctx context.Context, updatedVehicle data.VehicleStruct) (bool, error) {
	// Check if item exists
	oldVehicle, err := m.db.GetVehicle(updatedVehicle.VehicleId)
	if err != nil {
		logger.Error("err: ", err)
		return false, err
	}

	if len(oldVehicle) == 0 {
		return false, nil
	}

	// pass final item to DB
	err = m.db.UpdateVehicle(updatedVehicle)
	if err != nil {
		logger.Error("err: ", err)
		return false, err
	}

	return true, nil
}

func (m manager) ReserveVehicleForOrderForCity(ctx context.Context, city string) (vehicleId int, err error) {

	//1. Vehicle lock - activeCount + 1

	vehicleId, err = m.db.ReserveVehicleForOrderForCity(ctx, city)
	if err != nil {
		logger.Error("err: ", err)
		if xerrors.Is(err, utils.NoRowsFound) {
			return 0, NoVehicleFound
		}
		return 0, err
	}
	return vehicleId, nil
}

func (m manager) DecreaseActiveOrderCount(ctx context.Context, vehicleId int) (err error) {
	err = m.db.DecreaseActiveOrderCount(ctx, vehicleId)
	if err != nil {
		logger.Error("err: ", err)
		if xerrors.Is(err, utils.NorowsUpdated) {
			return CannotDecreaseCount
		}
		return err
	}
	return nil
}
