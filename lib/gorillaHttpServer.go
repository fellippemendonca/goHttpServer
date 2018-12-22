package lib;

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func GorillaHttpServer() {
	fmt.Println("\n\n ## The following content is exploring Gorilla Http Server ## \n");
	initServer();
}

func initServer() {
	r := mux.NewRouter();
	r.HandleFunc("/books/{title}", postBook).Methods("POST")
	r.HandleFunc("/books/{title}", getBook).Methods("GET")
	r.HandleFunc("/books/{title}/page/{page}", getBookPage).Methods("GET")
	r.HandleFunc("/books/{title}", putBook).Methods("PUT")
	r.HandleFunc("/books/{title}", deleteBook).Methods("DELETE")
	http.ListenAndServe(":8080", r);
}

func postBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	title := vars["title"];
	fmt.Fprintf(w, "You've requested POST book: %s\n", title);
}

func getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	title := vars["title"];
	fmt.Fprintf(w, "You've requested GET book: %s\n", title);
}

func getBookPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	title := vars["title"];
	page := vars["page"];
	fmt.Fprintf(w, "You've requested GET book: %s on page %s\n", title, page);
}

func putBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	title := vars["title"];
	fmt.Fprintf(w, "You've requested PUT book: %s\n", title);
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	title := vars["title"];
	fmt.Fprintf(w, "You've requested DELETE book: %s\n", title);
}