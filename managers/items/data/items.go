package data

import "database/sql"

type Item struct {
	ItemId int    `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}

type ItemForUpdate struct {
	ItemId int            `json:"id"`
	Name   sql.NullString `json:"name"`
	Price  sql.NullInt32  `json:"price"`
}
