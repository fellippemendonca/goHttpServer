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
	path := &Path{}
	path.uri = u
	path.head = nil
	path.mux = r.mux
	return path
}

// Init Router
func Init() *Router {
	r := Router{}
	r.mux = http.NewServeMux()
	return &r
}
