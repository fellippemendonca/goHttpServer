package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUsers(router *gin.RouterGroup) {
	fmt.Println("[OK] -- INITIALIZING USERS CONTROLLER")
	postUser(router)
	getUser(router)
	putUser(router)
	deleteUser(router)
}

func postUser(router *gin.RouterGroup) {
	router.POST("/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Hello %s", id)
	})
}

func getUser(router *gin.RouterGroup) {
	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		lastname := c.Query("lastname")
		firstname := c.DefaultQuery("firstname", "Guest")
		c.String(http.StatusOK, "Hello %s %s %s", id, firstname, lastname)
	})
}

func putUser(router *gin.RouterGroup) {
	router.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")
		lastname := c.Query("lastname")
		firstname := c.DefaultQuery("firstname", "Guest")
		c.String(http.StatusOK, "Hello %s %s %s", id, firstname, lastname)
	})
}

func deleteUser(router *gin.RouterGroup) {
	router.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		lastname := c.Query("lastname")
		firstname := c.DefaultQuery("firstname", "Guest")
		c.String(http.StatusOK, "Hello %s %s %s", id, firstname, lastname)
	})
}
