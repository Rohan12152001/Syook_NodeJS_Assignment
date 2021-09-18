package utils

import (
	"context"
	datas "github.com/Rohan12152001/Syook_Assignment/managers/user/data"
	"github.com/gin-gonic/gin"
)

var userKey = "user"

// For authMiddleWare
func SetContext(c *gin.Context, userObject datas.User) {
	c.Set(userKey, userObject)
	return
}

func GetUserFromContext(c context.Context) (datas.User, bool) {
	user := c.Value(userKey)
	if user == nil {
		return datas.User{}, false
	}
	user1, ok := user.(datas.User)
	if !ok {
		return datas.User{}, false
	}
	return user1, true
}
