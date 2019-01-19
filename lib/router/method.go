package router

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/fellippemendonca/goHttpServer/lib"
)

func NewMethod() *Method {
	m := &Method{}
	m.routes = make(map[*regexp.Regexp]http.Handler)
	return m
}

type Method struct {
	routes map[*regexp.Regexp]http.Handler
}

func (m *Method) Add(route string, next http.HandlerFunc) {
	m.routes[lib.GenerateRegex(route)] = next
}

func (m *Method) FindHandler(url string) http.Handler {
	for regex, handler := range m.routes {
		if regex.MatchString(url) {
			fmt.Println(regex)
			return handler
		}
	}
	return nil
}
