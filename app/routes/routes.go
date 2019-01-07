package routes

import (
	"fmt"
	//"github.com/fellippemendonca/goHttpServer/app/api"
	"github.com/fellippemendonca/goHttpServer/app/middleware"
	"net/http"
	"reflect"
)

func Init() *http.ServeMux {
	fmt.Println("[OK] -- INITIALIZING ROUTES")
	router := http.NewServeMux()
	fmt.Println(reflect.TypeOf(router))
	middleware.Init(router)
	//apiRouter := router.Group("/api")
	//api.Init(apiRouter)
	return router
}
