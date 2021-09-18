package user

import (
	"github.com/Rohan12152001/Syook_Assignment/managers/user/data"
)

type UserManager interface {
	GetUserbyId(id string) (datas.User, error)
	LoginUser(email, password string) (string, int, error)
}
