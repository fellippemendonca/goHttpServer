package routes

import (
	"fmt"
	"github.com/fellippemendonca/goHttpServer/app/api"
	"github.com/fellippemendonca/goHttpServer/app/middleware"
	"github.com/gin-gonic/gin"
	"reflect"
)

func Init() *gin.Engine {
	fmt.Println("[OK] -- INITIALIZING ROUTES")
	router := gin.Default()
	fmt.Println(reflect.TypeOf(router))
	middleware.Init(router)
	apiRouter := router.Group("/api")
	api.Init(apiRouter)
	return router
}
