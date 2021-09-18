package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

// This is not exactly Testing, I am still learning
func TestManager_CreateItem(t *testing.T) {
	db := New()
	id, err := db.CreateItem("rohan", 10001)
	fmt.Println(id, err)
}
