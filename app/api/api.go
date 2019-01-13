package api

import (
	"fmt"

	v1 "github.com/fellippemendonca/goHttpServer/app/api/v1"
	v2 "github.com/fellippemendonca/goHttpServer/app/api/v2"
	"github.com/fellippemendonca/goHttpServer/lib/router"
)

func Init(p *router.Path) {
	fmt.Println("[OK] -- INITIALIZING APIs")
	v1Path := p.Add("/v1")
	v2Path := p.Add("/v2")
	v1.Init(v1Path)
	v2.Init(v2Path)
}
