package v1

import (
	"fmt"
	"net/http"

	"github.com/fellippemendonca/goHttpServer/lib/router"
)

func InitUsers(p *router.Path) {
	fmt.Println("[OK] -- INITIALIZING USERS CONTROLLER")
	//users(p)
	usersById(p)
}
/*
func users(p *router.Path) {
	p.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("USERS"))
	})
}
*/

func usersById(p *router.Path) {
	p.Get("/:id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("USERS BY ID"))
	})
}
