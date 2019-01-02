package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {
	fmt.Println("[OK] -- INITIALIZING V2 CONTROLLERS")
	usersRouter := router.Group("/users")
	InitUsers(usersRouter)
}
