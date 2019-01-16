package router

import (
	"net/http"
)

// Router struct
type Router struct {
	mux *http.ServeMux
}

// Add a new path to the router
func (r *Router) Add(u string) *Path {
	path := NewPath()
	path.uri = u
	path.mux = r.mux
	return path
}

func (r *Router) GetMux() *http.ServeMux {
	return r.mux
}

// Init Router
func Init() *Router {
	r := Router{}
	r.mux = http.NewServeMux()
	return &r
}
