package v2

import (
	"fmt"
	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	fmt.Println("[OK] -- INITIALIZING V2 CONTROLLERS")
	usersRouter := router.PathPrefix("/users").Subrouter()
	InitUsers(usersRouter)
}
