package v2

import (
	"fmt"
	"net/http"

	"github.com/fellippemendonca/goHttpServer/lib/router"
)

func InitUsers(p *router.Path) {
	fmt.Println("[OK] -- INITIALIZING USERS CONTROLLER")
	h := p.Handler()
	usersGet(h)
	usersGetByID(h)
	usersPut(h)
	usersPost(h)
}

func usersGet(h *router.Handler) {
	h.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get users"))
	})
}

func usersGetByID(h *router.Handler) {
	h.Get("/:id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get users by ID"))
	})
}

func usersPost(h *router.Handler) {
	h.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Post users"))
	})
}

func usersPut(h *router.Handler) {
	h.Put("/:id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Put users by ID"))
	})
}
