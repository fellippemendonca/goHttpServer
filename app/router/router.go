package router

import (
	"fmt"
	"github.com/fellippemendonca/goHttpServer/app/controllers"
	"github.com/fellippemendonca/goHttpServer/app/middleware"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	fmt.Println("[OK] -- INITIALIZING ROUTES")
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	middleware.Init(router)
	controllers.Init(apiRouter)
	return router
}
