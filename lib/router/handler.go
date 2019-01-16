package router

import (
	"fmt"
	"net/http"
)

func NewHandler() *Handler {
	h := &Handler{}
	h.methods = make(map[string]*Method)
	return h
}

type Handler struct {
	route   string
	methods map[string]*Method
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method + "" + r.URL.Path)
	m := h.methods[r.Method]
	fmt.Println(h.methods)
	fmt.Println(m.routes)
	if m == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	handler := m.FindHandler(r.URL.Path)
	if handler == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	handler.ServeHTTP(w, r)
}

func (h *Handler) Get(s string, next http.HandlerFunc) {
	method := "GET"
	route := h.route + "" + s
	m := h.methods[method]
	if m == nil {
		m = NewMethod()
	}
	m.Add(route, next)
	h.methods[method] = m
}

func (h *Handler) Put(s string, next http.HandlerFunc) {
	method := "PUT"
	route := h.route + "" + s
	m := h.methods[method]
	if m == nil {
		m = NewMethod()
	}
	m.Add(route, next)
	h.methods[method] = m
}

func (h *Handler) Post(s string, next http.HandlerFunc) {
	method := "POST"
	route := h.route + "" + s
	m := h.methods[method]
	if m == nil {
		m = NewMethod()
	}
	m.Add(route, next)
	h.methods[method] = m
}
