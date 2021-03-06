package router

import (
	"fmt"
	"net/http"
)

// NewPath to be handled by the Mux handler
func NewPath() *Path {
	p := &Path{}
	p.head = nil
	return p
}

// Path struct
type Path struct {
	uri  string
	head *Path
	mux  *http.ServeMux
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
	// p.mux.Handle(h.route+"/", p.use(h))
	p.mux.Handle(h.route+"/", use(h))
	fmt.Println(h.route + "/")
	return h
}

func use(next http.Handler) http.Handler {
	fmt.Println("USE!!!")
	return next
}

/*
func (p *Path) use(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("URI:", r.RequestURI, ", Headers:", r.Header.Get("x-auth-token"))
		next.ServeHTTP(w, r)
	})
}
*/

func (p *Path) resolveRoute() string {
	if p.head != nil {
		return p.head.resolveRoute() + "" + p.uri
	}
	return p.uri
}
