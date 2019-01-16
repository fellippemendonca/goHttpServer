package router

import (
	"fmt"
	"github.com/fellippemendonca/goHttpServer/lib"
	"net/http"
	"regexp"
)

type Method struct {
	routes map[*regexp.Regexp]http.Handler
}

func NewMethod() *Method {
	m := &Method{}
	m.routes = make(map[*regexp.Regexp]http.Handler)
	return m
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
