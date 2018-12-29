package v1

import (
	"fmt"
	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	fmt.Println("[OK] -- INITIALIZING V1 CONTROLLERS")
	booksRouter := router.PathPrefix("/books").Subrouter()
	InitBooks(booksRouter)
	usersRouter := router.PathPrefix("/users").Subrouter()
	InitUsers(usersRouter)
}
