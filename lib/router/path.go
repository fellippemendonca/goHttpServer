package router

import (
	"net/http"
	"regexp"
)

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

func (p *Path) resolveRoute() string {
	if p.head != nil {
		return p.head.resolveRoute() + "" + p.uri
	}
	return p.uri
}

// Handler generator
func (p *Path) Handler() *Handler {
	h := &Handler{}
	h.route = p.resolveRoute()
	h.methods = make(map[*regexp.Regexp]http.Handler)
	p.mux.Handle(h.route+"/", h)
	return h
}
