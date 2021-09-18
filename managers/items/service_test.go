package items

// I was just trying my hands on unit testing

//import (
//	"context"
//	"database/sql"
//	"github.com/Rohan12152001/Syook_Assignment/managers/items/db"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestManager_CreateItem(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	mockDb := db.NewMockItemsDBManager(ctrl)
//	m := manager{db: mockDb}
//	mockDb.EXPECT().CreateItem("rohan", 10).Return(1, nil).Times(1)
//
//	id, err := m.CreateItem(context.Background(), "rohan", 10)
//	assert.NoError(t, err)
//	assert.EqualValues(t, 1, id)
//}
//
//func TestManager_CreateItem_Failure(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	mockDb := db.NewMockItemsDBManager(ctrl)
//	m := manager{db: mockDb}
//	mockDb.EXPECT().CreateItem("rohan", 10).Return(-1, sql.ErrNoRows).Times(1)
//
//	id, err := m.CreateItem(context.Background(), "rohan", 10)
//
//	assert.Error(t, err)
//	assert.EqualError(t, err, sql.ErrConnDone.Error())
//	assert.Equal(t, -1, id)
//}
