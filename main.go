package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Test
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	// r.Run()

}
