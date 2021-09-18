package db

import (
	"context"
	"fmt"
	"github.com/Rohan12152001/Syook_Assignment/managers/items/data"
	"github.com/Rohan12152001/Syook_Assignment/utils"
	"github.com/jmoiron/sqlx"
)

type manager struct {
	db *sqlx.DB
}

func New() ItemsDBManager {
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

func (m manager) GetAllItems() ([]data.Item, error) {
	query := "Select name,price from items"
	items := []data.Item{}

	err := m.db.Select(&items, query)
	if err != nil {
		return []data.Item{}, err
	}
	return items, nil
}

func (m manager) GetItem(Id int) (*data.Item, error) {
	query := "Select * from items where itemId=$1"
	items := []*data.Item{}

	err := m.db.Select(&items, query, Id)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, utils.NoRowsFound
	}

	return items[0], nil
}

func (m manager) CreateItem(ItemName string, ItemPrice int) (ItemId int, err error) {
	query := "Insert into items (name,price) values($1, $2) RETURNING itemId;"
	fmt.Println(query)

	Id := -1
	err = m.db.QueryRow(query, ItemName, ItemPrice).Scan(&Id)
	if err != nil {
		return -1, err
	}

	return Id, nil
}

func (m manager) UpdateItem(UpdatedItem data.Item) error {

	// Begin transaction
	tx, err := m.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}

	// Then create
	creatQuery := "UPDATE items SET name=$1, price=$2 where itemId=$3;"

	_, err = tx.Exec(creatQuery, UpdatedItem.Name, UpdatedItem.Price, UpdatedItem.ItemId)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
