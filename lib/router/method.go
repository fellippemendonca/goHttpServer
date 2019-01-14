package router

import (
	"fmt"
	"net/http"
)

func (p *Path) Get(h http.HandlerFunc) {
	uri, params := p.Resolve()
	fmt.Println(uri, params)
	p.router.mux.HandleFunc(uri, h)
}
