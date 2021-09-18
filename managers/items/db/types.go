package db

import (
	"github.com/Rohan12152001/Syook_Assignment/managers/items/data"
)

// go:generate mockgen -destination=mock_db.go -package=db -source=types.go
type ItemsDBManager interface {
	GetAllItems() ([]data.Item, error)
	GetItem(Id int) (*data.Item, error)
	CreateItem(ItemName string, ItemPrice int) (ItemId int, err error)
	UpdateItem(UpdatedItem data.Item) error
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "#yourPassword"
	dbname   = "Logistics"
)
