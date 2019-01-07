package middleware

import (
	"fmt"
	"github.com/fellippemendonca/goHttpServer/lib/logger"
	"github.com/fellippemendonca/goHttpServer/lib/auth"
	"net/http"
)

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

func Init(next http.Handler) http.Handler {
	fmt.Println("[OK] -- INITIALIZING MIDDLEWARE")
	return log(authenticate(next))
	//router.Handle("/", log(authenticate(next)))
}
