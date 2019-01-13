package router

import (
	"fmt"
	"net/http"
)

type MethodHandler struct{}

func (*MethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		fmt.Fprintf(w, "Method Allowed")
		return
	}
	fmt.Fprintf(w, "Method Not Allowed")
	return
	//fmt.Println("path", r.URL.Path)
	//query := r.URL.Query()
	//fmt.Println("query", query.Get("name"))
	//fmt.Fprintf(w, "Welcome to the users!")
}

func get(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		if r.Method == "GET" {
			fmt.Println("Method Allowed")
			f(w, r)
			return
		}
		fmt.Println("Method Not Allowed")
		return
	}
}

func (p *Path) Get(f http.HandlerFunc) {
	h := &MethodHandler{}
	uri, params := p.GetPath()
	fmt.Println(uri, params)
	//p.router.mux.HandleFunc(uri, get(f))
	p.router.mux.Handle(uri, h)
}

/**/
