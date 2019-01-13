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
}

func getUsers(p *router.Path) {
	path := p.Add("/")
	path.Get(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(p.GetUri())
		fmt.Println(p.GetParams())
		fmt.Fprintln(w, "Get Users")
	})
}

func getUserById(p *router.Path) {
	path := p.Add("/:id")
	path.Get(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Get Users By ID")
	})
}
