package app

import (
	"fmt"
	"github.com/fellippemendonca/goHttpServer/app/router"
	"net/http"
)

func Init() {
	fmt.Println("[OK] -- INITIALIZING HTTP SERVER")
	router := router.Init()
	http.ListenAndServe(":8080", router)
}
