package v2

import (
	"fmt"

	"github.com/fellippemendonca/goHttpServer/lib/router"
)

func Init(p *router.Path) {
	fmt.Println("[OK] -- INITIALIZING V2 CONTROLLERS")
	usersRouter := p.Add("/users")
	InitUsers(usersRouter)
}
