package db

import (
	datas "github.com/Rohan12152001/Syook_Assignment/managers/user/data"
)

type UserDb interface {
	GetUserbyID(Id string) (datas.User, error)
	GetUserForLogin(email, password string) (datas.User, error)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "#yourPassword"
	dbname   = "Logistics"
)
