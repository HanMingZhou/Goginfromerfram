package main

import (
	"Goginfromerfram/bootstrap"
	"Goginfromerfram/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	bootstrap.InitalizeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success")

	// Test
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	r.Run(":" + global.App.Config.App.Port)

}
