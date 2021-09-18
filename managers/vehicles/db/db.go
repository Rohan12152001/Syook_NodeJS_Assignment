package db

import (
	"context"
	_ "context"
	"database/sql"
	_ "errors"
	"fmt"
	// "github.com/Rohan12152001/Syook_Assignment/managers/vehicles"
	//"github.com/Rohan12152001/Syook_Assignment/managers/vehicles"
	"github.com/Rohan12152001/Syook_Assignment/managers/vehicles/data"
	"github.com/Rohan12152001/Syook_Assignment/utils"
	"github.com/jmoiron/sqlx"
	"golang.org/x/xerrors"
)

type manager struct {
	db *sqlx.DB
}

func New() VehiclesDBManager {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return manager{
		db: db,
	}
}

func (m manager) GetAllVehicles() ([]data.VehicleStruct, error) {
	query := "Select registrationNumber, vehicleType, city, activeOrdersCount from vehicles;"
	vehicles := []data.VehicleStruct{}

	err := m.db.Select(&vehicles, query)
	if err != nil {
		return []data.VehicleStruct{}, err
	}
	return vehicles, nil
}

func (m manager) GetVehicle(Id int) ([]data.VehicleStruct, error) {
	query := "Select * from vehicles where vehicleId=$1"
	vehicle := []data.VehicleStruct{}

	err := m.db.Select(&vehicle, query, Id)
	if err != nil {
		return []data.VehicleStruct{}, err
	}
	return vehicle, nil
}

func (m manager) CreateVehicle(newVehicle data.VehicleStruct) (int, error) {
	query := "Insert into vehicles (registrationNumber,vehicleType,city) values($1, $2, $3) RETURNING vehicleId;"
	fmt.Println(query)

	Id := -1
	err := m.db.QueryRow(query,
		newVehicle.RegistrationNumber,
		newVehicle.VehicleType, newVehicle.City).Scan(&Id)
	if err != nil {
		return -1, err
	}

	return Id, nil
}

func (m manager) UpdateVehicle(updatedVehicle data.VehicleStruct) error {
	// Begin transaction
	tx, err := m.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}

	// Then create
	creatQuery := "UPDATE vehicles SET registrationNumber=$1, vehicleType=$2, city=$3 where vehicleId=$4;"

	_, err = tx.Exec(
		creatQuery,
		updatedVehicle.RegistrationNumber,
		updatedVehicle.VehicleType,
		updatedVehicle.City,
		updatedVehicle.VehicleId)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (m manager) GetAvailableVehicleIds(matchCity string) (vehicles []data.VehicleStruct, err error) {
	query := "Select vehicleId, activeOrdersCount from vehicles where city=$1 and activeOrdersCount<=1;"
	vehicles = []data.VehicleStruct{}

	err = m.db.Select(&vehicles, query)
	if err != nil {
		return nil, err
	}

	return vehicles, nil

}

func (m manager) ReserveVehicleForOrderForCity(ctx context.Context, city string) (vehicleId int, err error) {

	txn, ok := utils.GetTransactionFromContext(ctx)
	if !ok {
		txn, err = m.db.BeginTx(ctx, nil)
		if err != nil {
			return -1, err
		}
		defer func() {
			if err == nil {
				txn.Rollback()
			} else {
				txn.Commit()
			}
		}()
	}

	var id int

	// Then update vehicles
	updateQuery := "update vehicles set activeOrdersCount = activeOrdersCount + 1 where vehicleId in (select vehicleId from vehicles where activeOrdersCount <= 1 and city = $1 limit 1 for update) returning vehicleId;"

	err = txn.QueryRow(updateQuery, city).Scan(&id)
	if err != nil {
		if xerrors.Is(err, sql.ErrNoRows) {
			return 0, utils.NoRowsFound
		}

		return -1, err
	}

	return id, nil
}

func (m manager) DecreaseActiveOrderCount(ctx context.Context, vehicleId int) (err error) {
	txn, ok := utils.GetTransactionFromContext(ctx)
	if !ok {
		txn, err = m.db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		defer func() {
			if err == nil {
				txn.Rollback()
			} else {
				txn.Commit()
			}
		}()
	}

	// Then update vehicles
	Id := -1
	updateQuery := "update vehicles set activeOrdersCount = activeOrdersCount - 1 where vehicleId=$1 and activeOrdersCount>0 returning vehicleId;"

	_ = txn.QueryRow(updateQuery, vehicleId).Scan(&Id)
	if err != nil {
		return err
	}

	if Id == -1 {
		return utils.NorowsUpdated
	}

	return nil
}
