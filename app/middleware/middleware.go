package middleware;

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	fmt.Println("\n\n ## INITIALYZING MIDDLEWARE ## \n");
	router.Use(middleware);
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		log.Println("URI:", r.RequestURI, "Headers:", r.Header)
		//fmt.Fprintf(w, "No no no. Wrong Path.");
	});
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-auth-token") != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			//fmt.Fprintf(w, "No no no. Wrong Pass.");
			return
		}
		log.Println("URI:", r.RequestURI, "Headers:", r.Header)
		next.ServeHTTP(w, r)
	})
}