package middleware

import (
	"fmt"
	//"github.com/fellippemendonca/goHttpServer/lib/logger"
	//"github.com/fellippemendonca/goHttpServer/lib/auth"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func Init(router *gin.Engine) {
	fmt.Println("[OK] -- INITIALIZING MIDDLEWARE")
	//router.Use(globalMiddleware("testString"))
	router.Use(log)
	authorized := router.Group("/")
	authorized.Use(authenticate)
	/*
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	*/
}

func globalMiddleware(params string) gin.HandlerFunc {
    // <---
	// This is part one
	fmt.Println("GLOBAL MIDDLEWARE ... 1")
    // --->
    // The follow code is an example
    //if err := check(params); err != nil {
        //panic(err)
    //}

    return func(c *gin.Context) {
        // <---
		// This is part two
		fmt.Println("GLOBAL MIDDLEWARE ... 2")
        // --->
        // The follow code is an example
        c.Set("TestVar", params)
		c.Next()
		fmt.Println("GLOBAL MIDDLEWARE ... 3")
    }
}

/**/
func authenticate(h gin.HandlerFunc) gin.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-auth-token")
		if auth.Check(token) != true {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func log(h gin.HandlerFunc) gin.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("URI:", r.RequestURI, ", Headers:", r.Header.Get("x-auth-token"))
	  	h.ServeHTTP(w, r)
	})
}
