package api

import (
	"fmt"
	"github.com/fellippemendonca/goHttpServer/app/api/v1"
	"github.com/fellippemendonca/goHttpServer/app/api/v2"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {
	fmt.Println("[OK] -- INITIALIZING APIs")
	v1Router := router.Group("/v1")
	v2Router := router.Group("/v2")
	v1.Init(v1Router)
	v2.Init(v2Router)
}
