package router

import (
	"net/http"
)

func Init() *Router {
	r := Router{}
	r.mux = http.NewServeMux()
	return &r
}

type Router struct {
	mux *http.ServeMux
}

func (r *Router) Add(u string) *Path {
	path := &Path{}
	path.uri = u
	path.head = nil
	path.mux = r.mux
	return path
}

func (r *Router) GetMux() *http.ServeMux {
	return r.mux
}
