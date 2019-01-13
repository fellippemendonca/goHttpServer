package routes

import (
	"fmt"
	"net/http"

	"github.com/fellippemendonca/goHttpServer/app/api"
	"github.com/fellippemendonca/goHttpServer/lib/router"
)

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// Router Initializer
func Init() *http.ServeMux {
	fmt.Println("[OK] -- INITIALIZING ROUTES")
	r := router.Init()
	//finalHandler := http.HandlerFunc(final)
	//handlers := middleware.Init(finalHandler)
	//router.Handle("/", handlers)
	apiPath := r.Add("/api")
	api.Init(apiPath)
	return r.GetMux()
}
