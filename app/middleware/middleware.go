package middleware

import (
	"fmt"
	"net/http"

	"github.com/fellippemendonca/goHttpServer/lib/auth"
	"github.com/fellippemendonca/goHttpServer/lib/logger"
)

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

// Middleware Initializer
func Init(next http.Handler) http.Handler {
	fmt.Println("[OK] -- INITIALIZING MIDDLEWARE")
	next = authenticate(next)
	next = log(next)
	return next
}
