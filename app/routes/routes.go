package routes

import (
	"fmt"
	//"github.com/fellippemendonca/goHttpServer/app/api"
	"github.com/fellippemendonca/goHttpServer/app/middleware"
	"net/http"
	//"reflect" fmt.Println(reflect.TypeOf(router))
)

func final(w http.ResponseWriter, r *http.Request) {
	//log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func Init() *http.ServeMux {
	fmt.Println("[OK] -- INITIALIZING ROUTES")
	router := http.NewServeMux()
	finalHandler := http.HandlerFunc(final)
	handlers := middleware.Init(finalHandler)
	router.Handle("/", handlers)
	//apiRouter := router.Group("/api")
	//api.Init(apiRouter)
	return router
}
