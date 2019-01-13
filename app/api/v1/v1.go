package v1

import (
	"fmt"

	"github.com/fellippemendonca/goHttpServer/lib/router"
)

func Init(p *router.Path) {
	fmt.Println("[OK] -- INITIALIZING V1 CONTROLLERS")
	usersRouter := p.Add("/users")
	InitUsers(usersRouter)
}
