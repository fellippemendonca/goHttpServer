package v1

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func InitBooks(router *mux.Router) {
	fmt.Println("[OK] -- INITIALIZING BOOKS CONTROLLER")
	postBook(router)
	getBook(router)
	getBookPage(router)
	putBook(router)
	deleteBook(router)
}

func postBook(router *mux.Router) {
	method := "POST"
	path := "/{title}"
	logic := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "You've requested the POST book: %s\n", title)
	}
	router.HandleFunc(path, logic).Methods(method)
}

func getBook(router *mux.Router) {
	method := "GET"
	path := "/{title}"
	logic := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "You've requested the NEW GET book: %s\n", title)
	}
	router.HandleFunc(path, logic).Methods(method)
}

func getBookPage(router *mux.Router) {
	method := "GET"
	path := "/{title}/page/{page}"
	logic := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the NEW POST book: %s, page: %s\n", title, page)
	}
	router.HandleFunc(path, logic).Methods(method)
}

func putBook(router *mux.Router) {
	method := "PUT"
	path := "/{title}"
	logic := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "You've requested the PUT book: %s\n", title)
	}
	router.HandleFunc(path, logic).Methods(method)
}

func deleteBook(router *mux.Router) {
	method := "DELETE"
	path := "/{title}"
	logic := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "You've requested the DELETE book: %s\n", title)
	}
	router.HandleFunc(path, logic).Methods(method)
}
