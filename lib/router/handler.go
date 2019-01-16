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

func (h *Handler) Get(path string, next http.HandlerFunc) {
	h.addMethod("GET", path, next)
}

func (h *Handler) Put(path string, next http.HandlerFunc) {
	h.addMethod("PUT", path, next)
}

func (h *Handler) Post(path string, next http.HandlerFunc) {
	h.addMethod("POST", path, next)
}

func (h *Handler) Delete(path string, next http.HandlerFunc) {
	h.addMethod("DELETE", path, next)
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

func (h *Handler) addMethod(name string, path string, next http.HandlerFunc) {
	route := h.route + "" + path
	m := h.methods[name]
	if m == nil {
		m = NewMethod()
	}
	m.Add(route, next)
	h.methods[name] = m
}
