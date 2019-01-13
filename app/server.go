package app

import (
	"fmt"
	"net/http"

	"github.com/fellippemendonca/goHttpServer/app/routes"
)

type usersHandler struct{}

func (usersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the users!")
}

type booksHandler struct{}

func (booksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the books!")
}

// Server Initializer
func Server() {
	fmt.Println("[OK] -- INITIALIZING HTTP SERVER")
	mux := routes.Init()
	http.ListenAndServe(":3000", mux)
}
