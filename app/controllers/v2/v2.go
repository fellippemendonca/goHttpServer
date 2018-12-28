package v2;

import (
	"fmt"
	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	fmt.Println("\n\n ## INITIALYZING V2 CONTROLLERS ## \n");
	usersRouter := router.PathPrefix("/users").Subrouter();
	InitUsers(usersRouter);
}