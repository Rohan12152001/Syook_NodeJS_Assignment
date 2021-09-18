package middleware

import (
	_ "github.com/Rohan12152001/Syook_Assignment/managers/user"
	"github.com/gin-gonic/gin"
)

type AuthManager interface {
	AuthMiddleWareWithUser(c *gin.Context)
}
