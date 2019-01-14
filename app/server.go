package app

import (
	"fmt"
	"net/http"
	"regexp"
	// "github.com/fellippemendonca/goHttpServer/app/routes"
)


type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
    switch {
    case rNum.MatchString(r.URL.Path):
        digits(w, r)
    default:
        w.Write([]byte("Unknown Pattern"))
    }
}

// Server Initializer
func Server() {
	fmt.Println("[OK] -- INITIALIZING HTTP SERVER")
	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})
	// mux.HandleFunc("/api/", route)
	// mux := routes.Init()
	http.ListenAndServe(":3000", mux)
}

var rNum = regexp.MustCompile(`/api/([^/]+)`)  // Has digit(s)
func digits(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Has Value"))
}
