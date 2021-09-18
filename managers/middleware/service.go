package middleware

import (
	"fmt"
	"github.com/Rohan12152001/Syook_Assignment/managers/user"
	"github.com/Rohan12152001/Syook_Assignment/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// secret key
var jwtkey = []byte("secret_key")

type AuthMiddleWare struct {
	userManager user.UserManager
}

func New() AuthManager {
	return AuthMiddleWare{
		userManager: user.New(),
	}
}

func (a AuthMiddleWare) AuthMiddleWareWithUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		var err error = nil
		token, err = c.Cookie("token")
		if err != nil {
			fmt.Println(err)
			c.Error(err)
			c.AbortWithStatus(401)
			return
		}
	}
	if token == "" {
		c.AbortWithStatus(401)
		return
	}

	tokenStr := token

	claims := &user.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtkey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println(err)
			c.Error(err)
			c.AbortWithStatus(401)
			return
		}
		fmt.Println(err)
		c.Error(err)
		c.AbortWithStatus(400)
		return
	}

	if !tkn.Valid {
		fmt.Println(err)
		c.Error(err)
		c.AbortWithStatus(401)
		return
	}

	UserId := claims.UserId
	userObject, err := a.userManager.GetUserbyId(UserId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(500)
		return
	}

	// Set in context & use ahead
	utils.SetContext(c, userObject)

	//user, ok := utils.GetUserFromContext(c)
	//if !ok {
	//	c.AbortWithStatus(500)
	//	return
	//}
	//
	//fmt.Println(user)
	// set context end

	c.Next()
}
