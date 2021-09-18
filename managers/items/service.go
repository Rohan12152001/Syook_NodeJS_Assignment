package items

import (
	"context"
	"github.com/Rohan12152001/Syook_Assignment/managers/items/data"
	"github.com/Rohan12152001/Syook_Assignment/managers/items/db"
	"github.com/Rohan12152001/Syook_Assignment/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type manager struct {
	db db.ItemsDBManager
}

func New() ItemsManager {
	return manager{
		db: db.New(),
	}
}

var logger = logrus.New()

func (m manager) GetAllItems(ctx context.Context) ([]data.Item, error) {
	// Call db
	items, err := m.db.GetAllItems()
	if err != nil {
		logger.Error("err: ", err)
		return nil, err
	}

	return items, nil
}

func (m manager) GetItem(ctx context.Context, Id int) (*data.Item, error) {
	// Call db
	item, err := m.db.GetItem(Id)
	if err != nil {
		if xerrors.Is(err, utils.NoRowsFound) {
			logger.Error("err: ", err)
			return nil, ItemNotFound
		}
		logger.Error("err: ", err)
		return nil, err
	}
	return item, nil
}

func (m manager) CreateItem(ctx context.Context, Name string, Price int) (int, error) {
	// Call db
	itemId, err := m.db.CreateItem(Name, Price)
	if err != nil {
		logger.Error("err: ", err)
		return -1, err
	}

	return itemId, nil
}

func (m manager) UpdateItem(ctx context.Context, updatedItem data.Item) (bool, error) {
	// Check if item exists
	_, err := m.db.GetItem(updatedItem.ItemId)
	if err != nil {
		logger.Error("err: ", err)
		return false, err
	}

	// pass final item to DB
	err = m.db.UpdateItem(updatedItem)
	if err != nil {
		logger.Error("err: ", err)
		return false, err
	}

	return true, nil
}
