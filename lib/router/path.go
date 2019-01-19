package router

import (
	"fmt"
	"net/http"

	"github.com/fellippemendonca/goHttpServer/lib/auth"
	"github.com/fellippemendonca/goHttpServer/lib/logger"
)

// NewPath to be handled by the Mux handler
func NewPath() *Path {
	p := &Path{}
	p.head = nil
	return p
}

// Path struct
type Path struct {
	uri        string
	middleware http.Handler
	head       *Path
	mux        *http.ServeMux
}

// Add a new Path
func (p *Path) Add(u string) *Path {
	path := &Path{}
	path.uri = u
	path.head = p
	path.mux = p.mux
	return path
}

// Handler generator
func (p *Path) NewHandler() *Handler {
	h := NewHandler()
	h.route = p.resolveRoute()
	p.mux.Handle(h.route+"/", mid(h))
	fmt.Println(h.route + "/")
	return h
}

func (p *Path) resolveRoute() string {
	if p.head != nil {
		return p.head.resolveRoute() + "" + p.uri
	}
	return p.uri
}

// Middleware Initializer
func mid(handler http.Handler) http.Handler {
	fmt.Println("[OK] -- INITIALIZING MIDDLEWARE")
	handler = authenticate(handler)
	handler = log(handler)
	return handler
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
