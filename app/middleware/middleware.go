package middleware

import (
	"fmt"
	"github.com/fellippemendonca/goHttpServer/lib/logger"
	"github.com/fellippemendonca/goHttpServer/lib/auth"
	"github.com/gorilla/mux"
	"net/http"
)

func Init(router *mux.Router) {
	fmt.Println("[OK] -- INITIALIZING MIDDLEWARE")
	router.Use(log)
	router.Use(authenticate)
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
}

func authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-auth-token")
		if auth.Check(token) != true {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("URI:", r.RequestURI, ", Headers:", r.Header.Get("x-auth-token"))
	  	h.ServeHTTP(w, r)
	})
}
