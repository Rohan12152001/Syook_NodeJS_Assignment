package items

import (
	"context"
	"fmt"
	"github.com/Rohan12152001/Syook_Assignment/managers/items/data"
)

var (
	ItemNotFound = fmt.Errorf("item not found")
)

type ItemsManager interface {
	GetAllItems(ctx context.Context) ([]data.Item, error)
	GetItem(ctx context.Context, Id int) (*data.Item, error)
	CreateItem(ctx context.Context, Name string, Price int) (int, error)
	UpdateItem(ctx context.Context, updatedItem data.Item) (bool, error)
}
