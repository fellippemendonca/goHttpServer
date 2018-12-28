package router

import (
	"fmt"
	"github.com/fellippemendonca/goHttpServer/app/middleware"
	"github.com/fellippemendonca/goHttpServer/app/controllers/v1"
	"github.com/fellippemendonca/goHttpServer/app/controllers/v2"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	fmt.Println("\n\n ## INITIALYZING ROUTES ## \n");
	router := mux.NewRouter();
	apiRouter := router.PathPrefix("/api").Subrouter();
	v1Router := apiRouter.PathPrefix("/v1").Subrouter();
	v2Router := apiRouter.PathPrefix("/v2").Subrouter();
	middleware.Init(router);
	v1.Init(v1Router);
	v2.Init(v2Router);
	return router;
}