package app

import (
	"fmt"
	"net/http"

	"github.com/fellippemendonca/goHttpServer/app/routes"
)

// Server Initializer
func Server() {
	fmt.Println("[OK] -- INITIALIZING HTTP SERVER")
	mux := routes.Init()
	http.ListenAndServe(":3000", mux)
}
