package main

import (
	"Goginfromerfram/bootstrap"
	"Goginfromerfram/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	bootstrap.InitalizeConfig()

	// Test
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	r.Run(":" + global.App.Config.App.Port)

}
