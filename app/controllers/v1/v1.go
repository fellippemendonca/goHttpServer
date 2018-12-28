package v1;

import (
	"fmt"
	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	fmt.Println("\n\n ## INITIALYZING V1 CONTROLLERS ## \n");
	booksRouter := router.PathPrefix("/books").Subrouter();
	InitBooks(booksRouter);
	usersRouter := router.PathPrefix("/users").Subrouter();
	InitUsers(usersRouter);
}
