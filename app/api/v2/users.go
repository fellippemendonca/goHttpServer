package v2

import (
	"fmt"
	"net/http"

	"github.com/fellippemendonca/goHttpServer/lib/router"
)

func InitUsers(p *router.Path) {
	fmt.Println("[OK] -- INITIALIZING USERS CONTROLLER")
	getUsers(p)
	getUserById(p)
	//postUser(p)
	//putUser(p)
	//deleteUser(p)
}

type usersHandler struct{}

func (usersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the users!")
}

func getUsers(p *router.Path) {
	p.Get("/", usersHandler{})
}

func getUserById(p *router.Path) {
	p.Get("/:id", usersHandler{})
}

/*
func postUser(p *router.Path) {
	p.Post("/:id", usersHandler{})
}

func putUser(p *router.Path) {
	p.Put("/:id", usersHandler{})
}

func deleteUser(p *router.Path) {
	p.Delete("/:id", usersHandler{})
}
*/
