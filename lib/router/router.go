package router

import (
	"net/http"
)

// Init Router
func NewRouter() *Router {
	r := Router{}
	r.mux = http.NewServeMux()
	return &r
}

// Router struct
type Router struct {
	mux *http.ServeMux
}

// Add a new path to the router
func (r *Router) NewPath(u string) *Path {
	path := NewPath()
	path.uri = u
	path.mux = r.mux
	return path
}

// GetMux from router
func (r *Router) GetMux() *http.ServeMux {
	return r.mux
}
