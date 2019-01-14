package routes

import (
	"fmt"
	"net/http"
	"regexp"
	//"github.com/fellippemendonca/goHttpServer/app/api"
	"github.com/fellippemendonca/goHttpServer/lib/router"
)

// Router Initializer
func Init() *http.ServeMux {
	fmt.Println("[OK] -- INITIALIZING ROUTES")
	r := router.Init()
	rootPath := r.Add("/")
	apiPath := rootPath.Add("api")
	middleware.Init(rootPath)
	api.Init(apiPath)
	return r.GetMux()
}
