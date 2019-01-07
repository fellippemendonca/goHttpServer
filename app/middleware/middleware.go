package middleware

import (
	"fmt"
	"github.com/fellippemendonca/goHttpServer/lib/logger"
	"github.com/fellippemendonca/goHttpServer/lib/auth"
	"net/http"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		//log.Println("Executing middlewareOne again")
	})
}
  
func middlewareTwo(next http.Handler) http.Handler {	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println("Executing middlewareTwo")
		if r.URL.Path != "/" {
			return
		}
		next.ServeHTTP(w, r)
		//log.Println("Executing middlewareTwo again")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	//log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("URI:", r.RequestURI, ", Headers:", r.Header.Get("x-auth-token"))
	  	next.ServeHTTP(w, r)
	})
}

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-auth-token")
		if auth.Check(token) != true {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Init(router *http.ServeMux) {
	fmt.Println("[OK] -- INITIALIZING MIDDLEWARE")
	//router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusNotFound)
	//})
	finalHandler := http.HandlerFunc(final)
	router.Handle("/", log(authenticate(finalHandler)))
}
