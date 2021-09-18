package auth

import (
	"encoding/json"
	"github.com/Rohan12152001/Syook_Assignment/managers/middleware"
	"github.com/Rohan12152001/Syook_Assignment/managers/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type Authendpoints struct {
	usermanager user.UserManager
	authManager middleware.AuthManager
}

type UserLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New() Authendpoints {
	return Authendpoints{
		usermanager: user.New(),
		authManager: middleware.New(),
	}
}

var logger = logrus.New()

func (A Authendpoints) SetRoutes(router *gin.Engine) {
	router.POST("/login", A.LoginHandler)
}

// LoginHandler for endpoint layer
func (A Authendpoints) LoginHandler(context *gin.Context) {
	var LoginPayload UserLoginPayload

	// b is bytes
	b, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		context.Error(err)
		context.Status(500)
		return
	}

	err = json.Unmarshal(b, &LoginPayload)
	if err != nil {
		logger.Error(err)
		context.Status(500)
		return
	}

	// jwt begin
	tokenString, expirationTime, err := A.usermanager.LoginUser(LoginPayload.Email, LoginPayload.Password)
	if err != nil {
		logger.Error(err)
		context.AbortWithStatus(500)
		return
	}

	context.SetCookie(
		"token",
		tokenString,
		expirationTime,
		"/",
		"localhost",
		false,
		false,
	)

}
