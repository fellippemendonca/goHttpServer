package app

import (
	"fmt"
	"github.com/fellippemendonca/goHttpServer/app/routes"
	"net/http"
)

/*
type server struct {
    db     *someDatabase
    router *someRouter
    email  EmailSender
}
*/

func Init() {
	fmt.Println("[OK] -- INITIALIZING HTTP SERVER")
	myMux := routes.Init()
	http.ListenAndServe(":8080", myMux)
}
