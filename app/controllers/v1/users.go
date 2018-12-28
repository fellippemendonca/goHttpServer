package v1;

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func InitUsers(router *mux.Router) {
	fmt.Println("\n\n ## INITIALYZING USERS CONTROLLER ## \n");
	postUser(router);
	getUser(router);
	putUser(router);
	deleteUser(router);
}

func postUser(router *mux.Router) {
	method := "POST";
	path := "/{id}";
	logic := func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r);
		id := vars["id"];
		fmt.Fprintf(w, "You've requested the POST user: %s\n", id);
	}
	router.HandleFunc(path, logic).Methods(method);
}

func getUser(router *mux.Router) {
	method := "GET";
	path := "/{id}";
	logic := func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r);
		id := vars["id"];
		fmt.Fprintf(w, "You've requested the GET user: %s\n", id);
	}
	router.HandleFunc(path, logic).Methods(method);
}

func putUser(router *mux.Router) {
	method := "PUT";
	path := "/{id}";
	logic := func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r);
		id := vars["id"];
		fmt.Fprintf(w, "You've requested the PUT user: %s\n", id);
	}
	router.HandleFunc(path, logic).Methods(method);
}

func deleteUser(router *mux.Router) {
	method := "DELETE";
	path := "/{id}";
	logic := func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r);
		id := vars["id"];
		fmt.Fprintf(w, "You've requested the DELETE user: %s\n", id);
	}
	router.HandleFunc(path, logic).Methods(method);
}