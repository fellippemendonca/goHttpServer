package v1

import (
	"fmt"
	"net/http"

	"github.com/fellippemendonca/goHttpServer/lib/router"
)

func InitUsers(p *router.Path) {
	fmt.Println("[OK] -- INITIALIZING USERS CONTROLLER")
	h := p.Handler()
	users(h)
	usersByID(h)
}

func users(h *router.Handler) {
	h.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("USERS"))
	})
}

func usersByID(h *router.Handler) {
	h.Get("/:id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("USERS BY ID"))
	})
}
