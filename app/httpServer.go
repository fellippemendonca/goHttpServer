package app;

import (
	"fmt"
	"net/http"
	"github.com/fellippemendonca/goHttpServer/app/router"
)

func Init() {
	fmt.Println("\n\n ## INITIALYZING HTTP SERVER ## \n");
	router := router.Init();
	http.ListenAndServe(":8080", router);
}