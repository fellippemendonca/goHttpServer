package app

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/fellippemendonca/goHttpServer/app/routes"
)

var rNum = regexp.MustCompile(`/api/v1/users/([^/]+)`) // Has digit(s)
func digits(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Has Value"))
}

type handler struct {
}

func (h *handler) apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	w.Write([]byte("Get All Users"))
}

func (h *handler) apiIDHandler(w http.ResponseWriter, r *http.Request) {
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
	//h := &handler{}
	fmt.Println("[OK] -- INITIALIZING HTTP SERVER")
	//mux := http.NewServeMux()
	//mux.HandleFunc("/api/v1/users", apiHandler)
	//mux.HandleFunc("/api/v1/users/123", apiIDHandler)
	mux := routes.Init()
	http.ListenAndServe(":3000", mux)
}
