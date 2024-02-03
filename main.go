package main

import (
	"Goginfromerfram/bootstrap"
	"Goginfromerfram/global"
)

// 后续了解一下viper zap功能包

func main() {

	bootstrap.InitalizeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据链接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// Test
	// r := gin.Default()
	// r.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "pong")
	// })

	// r.Run(":" + global.App.Config.App.Port)

	// 启动服务器
	bootstrap.RunServer()

}
