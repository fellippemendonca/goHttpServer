package controllers

import (
	"fmt"
	"github.com/fellippemendonca/goHttpServer/app/controllers/v1"
	"github.com/fellippemendonca/goHttpServer/app/controllers/v2"
	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	fmt.Println("[OK] -- INITIALIZING APIs")
	v1Router := router.PathPrefix("/v1").Subrouter()
	v2Router := router.PathPrefix("/v2").Subrouter()
	v1.Init(v1Router)
	v2.Init(v2Router)
}
